We store all data in `ecs.World` and build on top of it.\
We store path in `PathComponent` and upon retrival we store asset as interface in `CacheComponent`.\
We convert `PathComponent` to `CacheComponent` using dispatchers where each file extension has dedicated dispatcher.\
Using interfaces for `CacheComponent` doesn't affect performance heavily, its nothing in comparison to data stored and its processing.\
To release assets we just remove `CacheComponent` (recommended) or entity with this component.\
`CacheComponent` stores interface and we use\
`func GetAsset[Asset any](assets Service, assetID ecs.EntityID) (Asset, error)`\
to parse is to our asset type.
