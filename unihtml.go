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

// Package unihtml contains a plugin for the UniDoc.
package unihtml ;import (_df "bytes";_c "context";_cb "errors";_eg "fmt";_fc "github.com/unidoc/unihtml/client";_ce "github.com/unidoc/unihtml/content";_b "github.com/unidoc/unihtml/selector";_de "github.com/unidoc/unihtml/sizes";_fd "github.com/unidoc/unipdf/v3/common";
_dd "github.com/unidoc/unipdf/v3/common/license";_dc "github.com/unidoc/unipdf/v3/creator";_be "github.com/unidoc/unipdf/v3/model";_ec "github.com/unidoc/unipdf/v3/render";_d "image";_e "image/color";_f "math";_a "net/url";_gc "os";_fb "time";);

// WaitVisible waits for the provided selector to be ready.
// A selector might be  i.e.  `#example` for id  and `.example` for classes. The second parameter defines how to match given selector.
func (_fdd *Document )WaitVisible (sel string ,by ..._b .ByType ){_gfe :=_b .BySearch ;if len (by )> 0{_gfe =by [0];};_fdd ._dbf =append (_fdd ._dbf ,_fc .BySelector {Selector :sel ,By :_gfe });};

// SetPageHeight sets the page height for given document.
func (_dfg *Document )SetPageHeight (pageHeight _de .Length )error {_dfg ._gd =pageHeight ;_dfg ._ccg =_dc .PositionAbsolute ;return nil ;};var ErrContentNotDefined =_cb .New ("\u0068\u0074\u006d\u006c\u0020\u0064o\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0063\u006f\u006e\u0074\u0065\u006et\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");


// ConnectOptions creates UniHTML HTTP Client and tries to establish connection with the server.
func ConnectOptions (o Options )error {_fcd =_fc .New (_fc .Options {Hostname :o .Hostname ,Port :o .Port ,HTTPS :o .Secure });_dcc ,_cf :=_c .WithTimeout (_c .Background (),_fb .Second *5);defer _cf ();if _db :=_fcd .HealthCheck (_dcc );_db !=nil {return _db ;
};return nil ;};var ErrNoClient =_cb .New ("\u0055n\u0069\u0048\u0054\u004d\u004c\u0020\u0063\u006c\u0069\u0065\u006et\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");

// Document is HTML document wrapper that is used for extracting and converting HTML document into PDF pages.
type Document struct{_fg _ce .Content ;_fe margins ;_ccg _dc .Positioning ;_dbg ,_gf float64 ;_feb _de .PageSize ;_ggd ,_gd _de .Length ;_aa _de .Orientation ;_ea bool ;_ccgd _fb .Duration ;_ccb []_fc .BySelector ;_dbf []_fc .BySelector ;_fa *_fb .Duration ;
};

// NewDocument creates new HTML Document used as an input for the creator.Drawable.
func NewDocument (path string )(*Document ,error ){_ddc :=&Document {};_fac ,_cea :=_a .Parse (path );if _cea !=nil {return nil ,_cea ;};switch _fac .Scheme {case "\u0068\u0074\u0074\u0070","\u0068\u0074\u0074p\u0073":_ddc ._fg ,_cea =_ce .NewWebURL (path );
if _cea !=nil {return nil ,_cea ;};return _ddc ,nil ;};_gdg ,_cea :=_gc .Stat (path );if _cea !=nil {return nil ,_cea ;};if !_gdg .IsDir (){_ddc ._fg ,_cea =_ce .NewHTMLFile (path );}else {_ddc ._fg ,_cea =_ce .NewZipDirectory (path );};if _cea !=nil {return nil ,_cea ;
};return _ddc ,nil ;};

// SetMargins sets the Document Margins.
func (_dac *Document )SetMargins (left ,right ,top ,bottom float64 ){_dac ._fe .Left =_de .Point (left );_dac ._fe .Right =_de .Point (right );_dac ._fe .Top =_de .Point (top );_dac ._fe .Bottom =_de .Point (bottom );_dac ._ccg =_dc .PositionAbsolute ;
};

// SetPos sets the absolute position. Changes object positioning to absolute.
func (_ede *Document )SetPos (x ,y float64 ){_ede ._ccg =_dc .PositionAbsolute ;_ede ._dbg =x ;_ede ._gf =y ;};

// WaitReady waits for the provided selector to be ready.
// A selector might be  i.e.  `#example` for id  and `.example` for classes. The second parameter defines how to match given selector.
func (_gb *Document )WaitReady (sel string ,by ..._b .ByType ){_eeg :=_b .BySearch ;if len (by )> 0{_eeg =by [0];};_gb ._ccb =append (_gb ._ccb ,_fc .BySelector {Selector :sel ,By :_eeg });};

// SetMarginTop sets the left margin.
func (_ceaa *Document )SetMarginTop (margin _de .Length ){_ceaa ._fe .Top =margin };func (_acc *Document )getMargins ()margins {_gafg :=_acc ._fe ;if _acc ._ccg .IsRelative (){_gafg .Top =_de .Millimeter (1);_gafg .Left =_de .Millimeter (1);_gafg .Bottom =_de .Millimeter (1);
_gafg .Right =_de .Millimeter (1);return _gafg ;};if _gafg .Top ==nil {_gafg .Top =_de .Millimeter (10);};if _gafg .Bottom ==nil {_gafg .Bottom =_de .Millimeter (10);};if _gafg .Left ==nil {_gafg .Left =_de .Millimeter (10);};if _gafg .Right ==nil {_gafg .Right =_de .Millimeter (10);
};return _gafg ;};

// WriteToFile writes the document to a file defined by the output path.
func (_abd *Document )WriteToFile (outputPath string )error {if _daa :=_abd .validate ();_daa !=nil {return _daa ;};_dbbe :=_fb .Second *20+_abd ._ccgd ;_fgf ,_ade :=_c .WithTimeout (_c .Background (),_dbbe );defer _ade ();_aaf ,_ebc :=_abd .extract (_fgf ,_abd ._ggd ,_abd ._gd ,_abd .getMargins ());
if _ebc !=nil {return _ebc ;};_eaf :=_dc .New ();for _ ,_dec :=range _aaf {if _ebc =_eaf .AddPage (_dec );_ebc !=nil {return _ebc ;};};return _eaf .WriteToFile (outputPath );};

// SetLandscapeOrientation sets document landscape page orientation.
func (_adf *Document )SetLandscapeOrientation (){_adf ._aa =_de .Landscape };func _cae (_ga _d .Image )float64 {_faa :=_ga .Bounds ();var (_ed int ;_cge _e .Color ;_cgd ,_beg ,_daf uint32 ;);_cgc ,_bga :=_faa .Min .X ,_faa .Max .Y -1;_gaf :=_ga .At (_cgc ,_bga );
_bdg ,_faf ,_begb ,_ :=_gaf .RGBA ();_decb :=_bdg ==_f .MaxUint16 &&_faf ==_f .MaxUint16 &&_begb ==_f .MaxUint16 ;for _bga =_faa .Max .Y -1;_bga >=_faa .Min .Y ;_bga --{var _bf bool ;for _cgc =_faa .Min .X ;_cgc < _faa .Max .X ;_cgc ++{_cge =_ga .At (_cgc ,_bga );
_cgd ,_beg ,_daf ,_ =_cge .RGBA ();if (_decb &&!(_cgd ==_bdg &&_beg ==_faf &&_begb ==_daf ))||(!_decb &&(_f .Abs (float64 (_cgd )-float64 (_bdg ))/float64 (_f .MaxUint16 )> 0.03||_f .Abs (float64 (_beg )-float64 (_faf ))/float64 (_f .MaxUint16 )> 0.03||_f .Abs (float64 (_daf )-float64 (_begb ))/float64 (_f .MaxUint16 )> 0.03)){_bf =true ;
break ;};};if _bf {break ;};_ed =_bga ;};return float64 (_faa .Max .Y -_ed )/float64 (_faa .Max .Y );};type margins struct{Left ,Right ,Bottom ,Top _de .Length ;};

// ContainerComponent implements creator.containerElement interface.
func (_ddg *Document )ContainerComponent (container _dc .Drawable )(_dc .Drawable ,error ){switch container .(type ){case *_dc .Chapter :default:return nil ,_eg .Errorf ("\u0075\u006e\u0069\u0068t\u006d\u006c\u002e\u0044\u006f\u0063\u0075\u006d\u0065n\u0074\u0020\u0063\u0061\u006e\u0027\u0074\u0020\u0062\u0065\u0020\u0061\u0020\u0063\u006f\u006d\u0070\u006fn\u0065\u006e\u0074\u0020\u006ff\u0020\u0074\u0068\u0065\u0020\u0025\u0054\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",container );
};return _ddg ,nil ;};

// Options are the HTML Client options used for establishing the connection.
type Options struct{

// Hostname defines the hostname for the client. Default value: 'localhost'.
Hostname string ;

// Port defines the port at which the server works. Default value: '8080'
Port int ;

// Secure is the flag that states if the connection uses HTTPS protocol. Default Value: 'false'.
Secure bool ;

// Prefix is an option setting used when the server is working with the URI prefix. Default Value: ''.
Prefix string ;};var _ _dc .Drawable =(*Document )(nil );

// SetMarginBottom sets the left margin.
func (_begbg *Document )SetMarginBottom (margin _de .Length ){_begbg ._fe .Bottom =margin };

// TrimLastPageContent trims the last page content so that next creator blocks are located just at the end of given block.
func (_cac *Document )TrimLastPageContent (){_cac ._ea =true };

// SetMarginRight sets the left margin.
func (_fed *Document )SetMarginRight (margin _de .Length ){_fed ._fe .Right =margin };func (_eac *Document )validate ()error {if _fcd ==nil {return ErrNoClient ;};if _eac ._fg ==nil {return ErrContentNotDefined ;};return nil ;};var _fcd *_fc .Client ;

// Connect creates UniHTML HTTP Client and tries to establish connection with the server.
func Connect (path string )error {if _ef :=_gg ();_ef !=nil {return _ef ;};_ge ,_dg :=_fc .ParseOptions (path );if _dg !=nil {return _dg ;};_fcd =_fc .New (_ge );_ca ,_ab :=_c .WithTimeout (_c .Background (),_fb .Second *5);defer _ab ();if _eb :=_fcd .HealthCheck (_ca );
_eb !=nil {return _eb ;};return nil ;};

// SetPageSize sets the page default size.
func (_cba *Document )SetPageSize (pageSize _de .PageSize )error {if !pageSize .IsAPageSize (){return _cb .New ("\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0070\u0061\u0067\u0065\u0020s\u0069\u007a\u0065");
};_cba ._feb =pageSize ;_cba ._ccg =_dc .PositionAbsolute ;return nil ;};func (_cebe *Document )extract (_ccc _c .Context ,_bef ,_dga _de .Length ,_ggg margins )([]*_be .PdfPage ,error ){_daca :=_fc .BuildHTMLQuery ().SetContent (_cebe ._fg ).PageSize (_cebe ._feb ).PaperWidth (_bef ).PaperHeight (_dga ).Orientation (_cebe ._aa ).MarginLeft (_ggg .Left ).MarginRight (_ggg .Right ).MarginTop (_ggg .Top ).MarginBottom (_ggg .Bottom ).WaitTime (_cebe ._ccgd );
for _ ,_eea :=range _cebe ._ccb {_daca .WaitReady (_eea .Selector ,_eea .By );};for _ ,_cef :=range _cebe ._dbf {_daca .WaitVisible (_cef .Selector ,_cef .By );};_fdc ,_fca :=_daca .Query ();if _fca !=nil {return nil ,_fca ;};var _ae _c .CancelFunc ;if _cebe ._fa !=nil {_ccc ,_ae =_c .WithTimeout (_ccc ,*_cebe ._fa );
}else {_ccc ,_ae =_c .WithTimeout (_ccc ,_fb .Second *15);};defer _ae ();_ege ,_fca :=_fcd .ConvertHTML (_ccc ,_fdc );if _fca !=nil {return nil ,_fca ;};_bcb :=_df .NewReader (_ege .Data );_geb ,_fca :=_be .NewPdfReader (_bcb );if _fca !=nil {return nil ,_fca ;
};return _geb .PageList ,nil ;};

// SetTimeoutDuration sets the timeout duration
// the default timeout is 15 seconds.
func (_gbg *Document )SetTimeoutDuration (duration _fb .Duration ){_gbg ._fa =&duration };

// SetPageWidth sets the page width for given document.
func (_ag *Document )SetPageWidth (pageWidth _de .Length )error {_ag ._ggd =pageWidth ;_ag ._ccg =_dc .PositionAbsolute ;return nil ;};

// GetPdfPages is a function that converts provided input content and
func (_ee *Document )GetPdfPages (ctx _c .Context )([]*_be .PdfPage ,error ){if _cce :=_ee .validate ();_cce !=nil {return nil ,_cce ;};return _ee .extract (ctx ,_ee ._ggd ,_ee ._gd ,_ee .getMargins ());};func _gg ()error {_cc :=_dd .GetLicenseKey ();if _cc ==nil {return _cb .New ("\u006e\u006f\u0020\u006cic\u0065\u006e\u0073\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");
};if !_cc .IsLicensed (){return _cb .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u006f\u0072 \u006e\u006f\u0020\u006c\u0069\u0063\u0065n\u0073\u0065\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};return nil ;};

// GeneratePageBlocks implements creator.Drawable interface.
func (_dbb *Document )GeneratePageBlocks (ctx _dc .DrawContext )([]*_dc .Block ,_dc .DrawContext ,error ){if _ceb :=_dbb .validate ();_ceb !=nil {return nil ,ctx ,_ceb ;};var _cg []*_dc .Block ;_bg :=_dbb .getMargins ();_dcb ,_fbd :=_dbb ._ggd ,_dbb ._gd ;
if _dbb ._ccg .IsRelative (){_dcb ,_fbd =_de .Point (ctx .Width ),_de .Point (ctx .Height );ctx .X -=float64 (_bg .Left .Points ());};_fbf ,_fbe :=_dbb .extract (_c .Background (),_dcb ,_fbd ,_bg );if _fbe !=nil {return nil ,_dc .DrawContext {},_fbe ;};
for _bb ,_bd :=range _fbf {_dde ,_da :=_dc .NewBlockFromPage (_bd );if _da !=nil {return nil ,_dc .DrawContext {},_da ;};var _gdf float64 ;if _dbb ._ea &&_bb ==len (_fbf )-1{_def :=_ec .NewImageDevice ();_ddf ,_ddb :=_def .Render (_bd );if _ddb !=nil {return nil ,_dc .DrawContext {},_ddb ;
};_abg ,_ddb :=_bd .GetMediaBox ();if _ddb !=nil {return nil ,_dc .DrawContext {},_ddb ;};_ad :=_fb .Now ();_cgb :=_cae (_ddf );_gdf =_abg .Height ()*_cgb ;_fd .Log .Trace ("\u0054\u0072i\u006d\u006d\u0069\u006eg\u0020\u006ca\u0073\u0074\u0020\u0064\u006f\u0063\u0075\u006de\u006e\u0074\u0020\u0070\u0061\u0067\u0065\u0020\u0074\u0061\u006b\u0065n\u003a\u0020\u0025\u0076",_fb .Since (_ad ));
if _dbb ._fe .Bottom !=nil {_gdf -=float64 (_dbb ._fe .Bottom .Points ());};if _gdf < 0{_gdf =0;};_fd .Log .Trace ("C\u0072\u006f\u0070\u0070\u0069\u006e\u0067\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0027\u0073\u0020\u0070\u0061\u0067\u0065\u0020\u0025\u002e2\u0066 \u0070\u006f\u0069\u006et\u0073\u0020o\u0066\u0066\u0020\u0062\u006f\u0074\u0074\u006f\u006d\u0020\u006f\u0066\u0020\u006d\u0065\u0064\u0069\u0061\u0020\u0062\u006f\u0078\u000a",_gdf );
};_ba ,_fcc ,_da :=_dde .GeneratePageBlocks (ctx );if _da !=nil {return nil ,_dc .DrawContext {},_da ;};ctx =_fcc ;ctx .Y -=_gdf ;if _bb !=len (_fbf )-1&&ctx .Y > (ctx .PageHeight -ctx .Margins .Bottom )*.95{ctx .X =ctx .Margins .Left ;ctx .Y =ctx .Margins .Top ;
ctx .Page ++;};_cg =append (_cg ,_ba ...);};return _cg ,ctx ,nil ;};

// SetMarginLeft sets the left margin.
func (_fbdd *Document )SetMarginLeft (margin _de .Length ){_fbdd ._fe .Left =margin };

// WaitTime sets the waiting time before the webpage is rendered to PDF.
func (_cace *Document )WaitTime (duration _fb .Duration ){_cace ._ccgd =duration };