package foobar

/*
	@enum ModelTypeEnum::
		-> Standard = 1, Standard mesh type;
		-> SingleMesh = 4, Mesh consisting of bones and skin data;
		-> Morph;
		-> Sector, Holds objects;
*/

/*
	@enum MaterialFlags::
		->TEXTUREDIFFUSE = 0x00040000, whether diffuse texture is present;
		->COLORED = 0x08000000, whether to use diffuse color (only applies with diffuse texture);
		->MIPMAPPING = 0x00800000;
		->ANIMATEDTEXTUREDIFFUSE = 0x04000000;
		->ANIMATEXTEXTUREALPHA = 0x02000000;
		->DOUBLESIDEDMATERIAL = 0x10000000, whether backface culling should be off;
		->ENVIRONMENTMAP = 0x00080000, simulates glossy material with environment texture;
		->NORMALTEXTUREBLEND = 0x00000100, blend between diffuse and environment texture normally;
		->MULTIPLYTEXTUREBLEND = 0x00000200, blend between diffuse and environment texture by multiplying;
		->ADDITIVETEXTUREBLEND = 0x00000400, blend between diffuse and environment texture by addition;
		->CALCREFLECTTEXTUREY = 0x00001000;
		->PROJECTREFLECTTEXTUREY = 0x00002000;
		->PROJECTREFLECTTEXTUREZ = 0x00004000;
		->ADDITIONALEFFECT = 0x00008000, should be ALPHATEXTURE | COLORKEY | ADDITIVEMIXING;
		->ALPHATEXTURE = 0x40000000;
		->COLORKEY = 0x20000000;
		->ADDITIVEMIXING = 0x80000000, the object is blended against the world by adding RGB (see street lamps etc.);
*/

type Header struct {
	Magic           [4]int8 `spec:"plain"` /* Has to be 'PACK' */
	DirectoryOffset int32   /* Offset to the directory */
	DirectoryLength int32   /* Directory length */
}

type Directory struct {
	FileName     [56]int8 `spec:"string"` /* Archived file name */
	FilePosition int32    `json:"pos" spec:"cool"`
	FileLength   int32
}

type Model struct { /* Random Model structure */
	ModelType          ModelTypeEnum /* Model type */
	FaceGroupCount     uint16        /* Number of face groups */
	FaceGroups         []FaceGroup   /* model's face groups */
	MeshInstanceBuffer int32         `spec:"ptr"`
	TransformMatrix    [4][4]float32
	WorldMatrices      [12][16][16][4][32]uint8
}

type FaceGroup struct {
	MaterialID int32 /* -1 for default */
	FaceCount  uint16
	Faces      []int64 `spec:"plain"`
}
