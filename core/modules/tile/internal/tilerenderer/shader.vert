#version 450 core

out GS { flat int vertexID; }
vs_out;

void main() { vs_out.vertexID = gl_VertexID; }
