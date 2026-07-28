#version 450 core

layout(points) in;

layout(triangle_strip, max_vertices = 4) out;

//

in GS { flat int vertexID; }
gs_in[];

out FS {
  vec2 uv;
  flat int textures[4];
}
gs_out;

layout(std430, binding = 0) buffer Grid { int grid[]; };

layout(std430, binding = 1) buffer TileTextures {
  vec2 tileTextures[]; // [index, amount]
};

uniform uint size;

vec2 getCoord(int i) {
  return vec2(        // this uses +1 because of dual grid system
      i % (size + 1), // X: Coord(index % g.width)
      i / (size + 1)  // Y: Coord(index >> g.width)
  );
}

int getIndex(vec2 coord) {
  coord = clamp(coord, vec2(0), vec2(size + 1));
  return int(coord.x) + int(coord.y) * int(size + 2);
}

int getTile(vec2 coord) { return grid[getIndex(coord)]; }
int getTile(int index) { return getTile(getCoord(index)); }

int getTexture(int seed, int i) {
  return int(tileTextures[i].x) + seed % int(tileTextures[i].y);
}

int seed(int num) {
  uint seed = num;
  seed = seed * 1664525u + 1013904223u;
  seed ^= seed >> 16;
  seed *= 0x85ebca6bu;
  seed ^= seed >> 13;
  seed *= 0xc2b2ae35u;
  seed ^= seed >> 16;
  return int(seed);
}

//

void sort4(inout int a[4]) {
  int tmp;
#define SWAP(i, j)                                                             \
  if (a[i] < a[j]) {                                                           \
    tmp = a[i];                                                                \
    a[i] = a[j];                                                               \
    a[j] = tmp;                                                                \
  }
  SWAP(0, 1);
  SWAP(2, 3);
  SWAP(0, 2);
  SWAP(1, 3);
  SWAP(1, 2);
#undef SWAP
}

void setBiomes(int index, vec2 coord) {
  int neighbours[4] = {getTile(coord + vec2(0)), getTile(coord + vec2(1, 0)),
                       getTile(coord + vec2(0, 1)), getTile(coord + vec2(1))};
  int biomes[4] = neighbours;
  sort4(biomes);

  int seed = seed(index);
  for (int i = 0; i < 4; i++) {
    int base = biomes[i] * 15;
    int mask = 0;
    for (int n = 0; n < 4; n++) {
      mask |= int(neighbours[n] == biomes[i]) * (1 << n);
    }
    gs_out.textures[i] = getTexture(seed, base + mask);
  }
}

uniform mat4 mvp;
uniform float sizeInv; // = 2 / size

vec2 corners[4] =
    vec2[](vec2(-.5, -.5), vec2(-.5, 0.5), vec2(0.5, -.5), vec2(0.5, 0.5));

void main() {
  int i = gs_in[0].vertexID;

  vec2 coord = getCoord(i);
  setBiomes(i, coord);

  for (int i = 0; i < 4; i++) {
    vec2 rawOffset = coord + corners[i];

    vec2 offset = clamp(rawOffset, vec2(0.0), vec2(size));
    vec2 uv = corners[i] + vec2(0.5);
    uv += step(rawOffset, vec2(0.0)) * 0.5;
    uv -= step(vec2(size), rawOffset) * 0.5;

    vec4 pos = vec4(sizeInv * offset - vec2(1.0), 0.0, 1.0);
    gl_Position = mvp * pos;
    gs_out.uv = uv;
    EmitVertex();
  }

  EndPrimitive();
}
