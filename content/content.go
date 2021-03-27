//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package content ;import (_b "archive/zip";_g "bytes";_ce "encoding/json";_cgc "golang.org/x/xerrors";_e "io/ioutil";_cg "net/url";_dc "os";_d "strings";);func (_bgd *zipDirectory )zipPath (_ef ,_db string )error {_bc ,_ae :=_e .ReadDir (_ef );if _ae !=nil {return _cgc .Errorf ("\u0072\u0065\u0061di\u006e\u0067\u0020\u0064\u0069\u0072\u0065\u0063\u0074o\u0072y\u003a \u0027%\u0073\u0027\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0077",_ef ,_ae );
};_bcf :=&_d .Builder {};for _ ,_gb :=range _bc {_bb ,_cgd :=_bgd .zipBasePaths (_gb ,_ef ,_db ,_gb .IsDir (),_bcf );if _gb .IsDir (){if _cbc :=_bgd .zipPath (_bb ,_cgd );_cbc !=nil {return _cbc ;};continue ;};_dg ,_ee :=_e .ReadFile (_bb );if _ee !=nil {return _ee ;
};_bce ,_ee :=_bgd ._cgae .Create (_cgd );if _ee !=nil {return _ee ;};if _ ,_ee =_bce .Write (_dg );_ee !=nil {return _ee ;};};return nil ;};

// NewHTMLFile creates new Content htmFile for provided input path.
func NewHTMLFile (path string )(Content ,error ){_f ,_ed :=_dc .Open (path );if _ed !=nil {return nil ,_ed ;};_gc :=&htmFile {_cb :_g .Buffer {}};if _ ,_ed =_gc ._cb .ReadFrom (_f );_ed !=nil {return nil ,_ed ;};return _gc ,nil ;};

// ContentType implements Content interface.
func (_bg *htmFile )ContentType ()string {return "\u0074e\u0078\u0074\u002f\u0068\u0074\u006dl"};

// Data implements Content interface.
func (_cbb *htmFile )Data ()[]byte {return _cbb ._cb .Bytes ()};

// Data implements Content interface.
func (_ba *webURL )Data ()[]byte {return _ba ._gd };type webURL struct{_a string ;_gd []byte ;};func (_dgf *zipDirectory )zipBasePaths (_ga _dc .FileInfo ,_fac ,_gbg string ,_fag bool ,_cbf *_d .Builder )(string ,string ){_cbf .WriteString (_fac );if !_d .HasSuffix (_fac ,"\u002f"){_cbf .WriteRune ('/');
};_cbf .WriteString (_ga .Name ());if _fag {_cbf .WriteRune ('/');};_de :=_cbf .String ();_cbf .Reset ();_cbf .WriteString (_gbg );_cbf .WriteString (_ga .Name ());if _fag {_cbf .WriteRune ('/');};_eg :=_cbf .String ();_cbf .Reset ();return _de ,_eg ;};


// NewWebURL creates new Content webURL for provided input URL path.
func NewWebURL (path string )(Content ,error ){if _ ,_bf :=_cg .Parse (path );_bf !=nil {return nil ,_bf ;};type urlOutput struct{URL string `json:"url"`;};_da :=urlOutput {URL :path };_be ,_bd :=_ce .Marshal (_da );if _bd !=nil {return nil ,_bd ;};return &webURL {_a :path ,_gd :_be },nil ;
};

// NewZipDirectory creates new zip compressed file that recursively reads the directory at the 'dirPath'
// and stores in it's in-memory buffer.
func NewZipDirectory (dirPath string )(Content ,error ){_eb :=&zipDirectory {_cga :_g .Buffer {}};_eb ._cgae =_b .NewWriter (&_eb ._cga );if _fed :=_eb .zipPath (dirPath ,"");_fed !=nil {return nil ,_fed ;};if _dd :=_eb ._cgae .Close ();_dd !=nil {return nil ,_dd ;
};return _eb ,nil ;};

// Content is an interface used for putting the content into Client Query.
type Content interface{ContentType ()string ;Data ()[]byte ;};

// ContentType implements Content interface.
func (_df *zipDirectory )ContentType ()string {return "\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u007a\u0069\u0070";};type zipDirectory struct{_cga _g .Buffer ;_cgae *_b .Writer ;};type htmFile struct{_cb _g .Buffer };

// Data implements Content interface.
func (_fa *zipDirectory )Data ()[]byte {return _fa ._cga .Bytes ()};

// ContentType implements Content interface.
func (_bgc *webURL )ContentType ()string {return "\u0061\u0070p\u006c\u0069\u0063a\u0074\u0069\u006f\u006e\u002f\u006a\u0073\u006f\u006e";};