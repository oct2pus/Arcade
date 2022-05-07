Okay so I had to stop this model prematurely, I noticed the hole for the heat set insert was fairly shorter than it should be (it should be approximately 12.7mm, i think it was 7mm).

To model I was using the [deadsy/sdfx](https://github.com/deadsy/sdfx) golang library, because I'm not right in the head and want to do code cad. I've been finding translate hard but I decided to figure out why things never seem to move in a predictable way.

Models start centered at `{0,0,0}`. i need to take the `difference` of the `insert` from the `sphere`, so if i do
```go
sphere = sdf.Difference3D(sphere,insert)
``` 
my hole is in the middle of the model. before i can do that I have to translate the insert on the Z (up/down) axis.

so i wrote 
```go
shift := sphere.BoundingBox.Max.Z - insert.BoundingBox.Max.Z / 2.0
insert = sdf.Transform3D(insert, sdf.Translate3d(sdf.V3{X: 0, Y: 0, Z: -shift}))
sphere = sdf.Difference3D(sphere, insert)
```
cool, looks good. let's print.

then while it was printing i noticed on the slicer that the layers did not add up; every layer is 0.16mm, `insert` is 13mm, i needed around 81 layers of hole for it to be approximately correct, and i was well short of that.

oh okay, i look again and noticed I forgot to **P**lease **E**xcuse **M**y **D**ear **A**unt **S**ally. Alright.

```go
shift := (sphere.BoundingBox.Max.Z - insert.BoundingBox.Max.Z) / 2.0

