# TODO
This list contains a list of tasks to keep in mind. Often architectural changes to revise and either implement it or omit it entirely
## [cicd](/cicd/readme/README.md)
Ensure script wraps itself in `Dockerfile` if it runs on local machine
- ### [docs](/cicd/modules/docs/readme/README.md)
1. Check is there better readme extension than ".md".
Extension to preview:
- ".rst"
- ".adoc"

2. Re work module interface to take struct and output string
## [core](/core/readme/README.md)
- ### [definitions](/core/modules/definitions/readme/README.md)
Move.
Clean up `/assets` and start using it instead of local copy.

Create automatic export from `.pxo` to `.png` and/or `.git`
- ### [pathfind](/core/modules/pathfind/readme/README.md)
Modify speed to allow moving many tiles per tick.
Modify pathfinding to do not use shortest route and instead use optimal path (optimize paths per chunks).

Research materials:
- `supreme commander`
- `multi agent pathfinding`
- ### [player](/core/modules/player/readme/README.md)
Restrict actions to allow only user to perform his actions.
Perhaps attach `PlayerComponent` to camera.
## [engine](/engine/readme/README.md)
- ### [camera](/engine/modules/camera/readme/README.md)
clean up pkg and move logic to internal
- ### [collider](/engine/modules/collider/readme/README.md)
Implement `CollidesWithObject`

Change main algorithm from spatial algorithm to tree algorithm.
Create methods to perform only shallow comparisons or only deep comparisons.
- ### [graphics](/engine/modules/graphics/readme/README.md)
Check `vulcan` and `wasm` as opengl alternatives
- ### [hierarchy](/engine/modules/hierarchy/readme/README.md)
abstract `InheritGroups` to inherit allowing to inherit any component.
Try to reduce calls during removal.
- ### [inputs](/engine/modules/inputs/readme/README.md)
Implement a proper input cursor and improve focusing and unfocusing on input
- ### [netsync](/engine/modules/netsync/readme/README.md)
Create more features to allow more specific features to allow more specific calls
- ### [noise](/engine/modules/noise/readme/README.md)
Improve noise flatenning bell curve.
Avoid equasions for flattening and instead create 10k sample slice to create a bell curve.
- ### [render](/engine/modules/render/readme/README.md)
Add frustrum culling. For each camera store all colliding objects.
- ### [text](/engine/modules/text/readme/README.md)
1. Implement more letters and make them easier to expand (emojis, hiragana, etc.).
2. Write own glyph rendering (solves first issue).
- ### [transform](/engine/modules/transform/readme/README.md)
Start using Fixed-Point math
