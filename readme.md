Demo Time
=========

This is a demo game, showing how to do these in Go on Windows:

- 3D graphics with Direct3D9
- Custom audio mixer with DirectSound8
- XBox controller input with XInput
- Joystick input with DirectInput
- Wavefront OBJ 3D model loading
- Load MP3 and OGG files

3D Modelling
============

We open the 3D models in Blender and scale them by -1 in Blender's Y-direction,
which in D3D is the Z-direction.

This way, the default Wavefront .obj export with Forward Axis -Z and Up Axis Y
will be right for D3D, which uses a left-handed coordinate system where
positive Z goes into the monitor, negative Z goes out of the monitor towards
the user.

We also flip the textures upside-down because Blenders UVs are like OpenGL,
they go from the bottom up, while D3D has its UVs go from top to bottom.

Credits
=======

These assets were used in the program:
-	[XBox 360 Controller](https://sketchfab.com/3d-models/xbox-360-controller-08bb78e6d7344a5cbf339123bd138966)
-	[Lego Joker](https://sketchfab.com/3d-models/joker-lego-1a998ccdfe3442ffa327b16637eb8032)
-	[Joystick](https://sketchfab.com/3d-models/joystick-3630d34457bd4de08db183cb4b106be9)
-	[Lego Brick](https://poly.pizza/m/f4-AUM_gO-R)
-	[Background music](https://pixabay.com/music/synthwave-stranger-things-124008/)
-	[Foot step sound](https://pixabay.com/sound-effects/st3-footstep-sfx-323056/)
-	The level texture was taken from Google Images, but I forgot to write down
	the actual links when composing the texture. Please don't sue.
