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

package cmd ;import (_b "context";_ea "fmt";_fg "github.com/mitchellh/go-homedir";_a "github.com/spf13/cobra";_eaf "github.com/spf13/viper";_ad "github.com/unidoc/unihtml/client";_aee "github.com/unidoc/unihtml/content";_d "github.com/unidoc/unihtml/sizes";
_ae "github.com/unidoc/unipdf/v3/common";_e "os";_ga "path/filepath";_f "time";);type generateConfig struct{Port int `mapstructure:"port"`;Host string `mapstructure:"host"`;Https bool `mapstructure:"https"`;Prefix string `mapstructure:"prefix"`;};

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute (){if _gc :=_gaf .Execute ();_gc !=nil {_ea .Println (_gc );_e .Exit (1);};};func _fa (){if _fe !=""{_eaf .SetConfigFile (_fe );}else {_bbg ,_dc :=_fg .Dir ();if _dc !=nil {_ea .Println (_dc );_e .Exit (1);};_eaf .AddConfigPath (_bbg );_eaf .SetConfigName ("\u002e\u0075\u006ei\u0068\u0074\u006d\u006c\u002d\u0073\u0072\u0063");
};_eaf .AutomaticEnv ();if _bae :=_eaf .ReadInConfig ();_bae ==nil {_ea .Println ("\u0055s\u0069n\u0067\u0020\u0063\u006f\u006ef\u0069\u0067 \u0066\u0069\u006c\u0065\u003a",_eaf .ConfigFileUsed ());};};func _bc (cmd *_a .Command ,_da []string ){_gf :=_f .Now ();
if _gfe :=_eaf .BindPFlags (cmd .Flags ());_gfe !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_gfe );_e .Exit (1);};if _gg :=_eaf .Unmarshal (&_c );_gg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_gg );
_e .Exit (1);};_ff ();_bga :=_f .Now ();_dd ,_ddg :=_e .Stat (_da [0]);if _ddg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ddg );_e .Exit (1);};if !_dd .IsDir (){if _ga .Ext (_dd .Name ())!="\u002e\u0068\u0074m\u006c"{_ea .Printf ("\u0045\u0072r\u003a\u0020\u0043\u0075\u0072\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006f\u006e\u006c\u0079\u0020\u0048\u0054M\u004c\u0020\u0066\u0069\u006c\u0065s\u0020\u0069\u006e\u0070\u0075\u0074\u0053\u0074\u0061\u0074\u0020\u0061\u0072\u0065 \u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u002e\u0020\u0049\u006ep\u0075\u0074\u003a\u0020\u0025\u0073\u000a",_da [0]);
_e .Exit (1);};};_gfb ,_ddg :=_e .OpenFile (_da [1],_e .O_CREATE |_e .O_WRONLY |_e .O_TRUNC ,0700);if _ddg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ddg );_e .Exit (1);};defer _gfb .Close ();_be :=_ad .New (_ad .Options {HTTPS :_c .Https ,Hostname :_c .Host ,Port :_c .Port ,Prefix :_c .Prefix });
_eeb ,_ab :=_b .WithTimeout (_b .Background (),_f .Second *10);defer _ab ();_bga =_f .Now ();var _bcf _aee .Content ;if _dd .IsDir (){_bcf ,_ddg =_aee .NewZipDirectory (_da [0]);}else {_bcf ,_ddg =_aee .NewHTMLFile (_da [0]);};if _ddg !=nil {_ea .Printf ("\u0045r\u0072\u003a\u0020\u0025\u0076",_ddg );
_e .Exit (1);};_ba ,_ddg :=_ad .BuildHTMLQuery ().PaperWidth (_eafc .PaperWidth .Length ).PaperHeight (_eafc .PaperHeight .Length ).PageSize (_eafc .PageSize ).MarginTop (_eafc .MarginTop .Length ).MarginBottom (_eafc .MarginBottom .Length ).MarginLeft (_eafc .MarginLeft .Length ).MarginRight (_eafc .MarginRight .Length ).Orientation (_eafc .Orientation ).SetContent (_bcf ).Query ();
if _ddg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ddg );_e .Exit (1);};_gag ,_ddg :=_be .ConvertHTML (_eeb ,_ba );if _ddg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ddg );_e .Exit (1);};_ae .Log .Trace ("\u0045\u0078\u0065cu\u0074\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0065r\u0061t\u0065 \u0071u\u0065\u0072\u0079\u0020\u0074\u0061\u006b\u0065\u006e\u003a\u0020\u0025\u0073",_f .Since (_bga ));
_bga =_f .Now ();_ ,_ddg =_gfb .Write (_gag .Data );if _ddg !=nil {_ea .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ddg );_e .Exit (1);};_ae .Log .Trace ("\u0057\u0072\u0069\u0074in\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0061\u006b\u0065\u006e\u003a\u0020%\u0073",_f .Since (_bga ));
_ea .Printf ("\u0047\u0065n\u0065\u0072\u0061\u0074\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0020\u0069\u006e %\u0073\u000a",_f .Since (_gf ));};func init (){_gaf .AddCommand (_bb );_bb .Flags ().IntP ("\u0070\u006f\u0072\u0074","\u0070",8080,"\u0050\u006f\u0072\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065 \u0075\u006e\u0069\u0068\u0074\u006d\u006c\u0020\u0073\u0065r\u0076\u0065\u0072");
_bb .Flags ().String ("\u0068\u006f\u0073\u0074","\u006co\u0063\u0061\u006c\u0068\u006f\u0073t","\u0048\u006f\u0073t\u0020\u006e\u0061\u006de\u0020\u006f\u0066\u0020\u0074\u0068\u0065 \u0075\u006e\u0069\u0068\u0074\u006d\u006c\u0020\u0073\u0065\u0072\u0076\u0065\u0072");
_bb .Flags ().BoolP ("\u0068\u0074\u0074p\u0073","\u0073",false ,"\u0050\u0072o\u0074\u006f\u0063\u006fl\u0020\u0075s\u0065\u0064\u0020\u0069\u006e\u0020\u0073\u0065r\u0076\u0065\u0072\u0020\u0063\u006f\u006d\u006d\u0075\u006e\u0069\u0063a\u0074\u0069\u006f\u006e");
_bb .Flags ().StringP ("\u0070\u0072\u0065\u0066\u0069\u0078","\u0078","","\u0050u\u0062\u006ci\u0063\u0020\u0061\u0070i\u0020\u0070\u0072e\u0066\u0069\u0078\u0020\u0075\u0073\u0065\u0064\u0020by\u0020\u0074\u0068e\u0020\u0075n\u0069\u0068\u0074\u006d\u006c\u0020s\u0065\u0072v\u0065\u0072");
_bb .Flags ().Var (&_eafc .PaperWidth ,"p\u0061\u0070\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068","\u0073\u0065\u0074s \u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0070\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068");_bb .Flags ().Var (&_eafc .PaperHeight ,"\u0070\u0061\u0070e\u0072\u002d\u0068\u0065\u0069\u0067\u0068\u0074","\u0073e\u0074\u0073\u0020\u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070a\u0070\u0065\u0072\u002d\u0068\u0065\u0069\u0067\u0068\u0074");
_bb .Flags ().Var (&_eafc .PageSize ,"\u0070\u0061\u0070\u0065\u0072\u002d\u0073\u0069\u007a\u0065","s\u0065\u0074\u0073\u0020up\u0020t\u0068\u0065\u0020\u0070\u0061g\u0065\u0020\u0073\u0069\u007a\u0065");_bb .Flags ().Var (&_eafc .Orientation ,"o\u0072\u0069\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e","\u0073\u0065\u0074\u0073 \u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067e\u0020o\u0072\u0069\u0065\u006e\u0074\u0061\u0074i\u006f\u006e");
_bb .Flags ().Var (&_eafc .MarginTop ,"\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070","\u0073\u0065\u0074\u0073 u\u0070\u0020\u0074\u0068\u0065\u0020\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074o\u0070");_bb .Flags ().Var (&_eafc .MarginBottom ,"\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0062\u006f\u0074\u0074\u006f\u006d","\u0073e\u0074\u0073\u0020\u0075p\u0020\u0074\u0068\u0065\u0020m\u0061r\u0067i\u006e\u002d\u0062\u006f\u0074\u0074\u006fm");
_bb .Flags ().Var (&_eafc .MarginRight ,"\u006d\u0061\u0072g\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074","\u0073e\u0074\u0073\u0020\u0075\u0070\u0020\u0074\u0068\u0065\u0020\u006da\u0072\u0067\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074");_bb .Flags ().Var (&_eafc .MarginLeft ,"m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074","\u0073\u0065\u0074\u0073 u\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0070\u0065\u0072\u002d\u006c\u0065f\u0074");
};var _gaf =&_a .Command {Use :"\u0075n\u0069\u0068\u0074\u006d\u006c",Short :"\u0041\u0020\u0062\u0072\u0069\u0065\u0066\u0020\u0064\u0065\u0073\u0063\u0072i\u0070\u0074\u0069\u006f\u006e\u0020o\u0066\u0020\u0079\u006f\u0075\u0072\u0020\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006f\u006e",Long :"\u0041\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0064\u0065\u0073c\u0072\u0069\u0070\u0074i\u006f\u006e\u0020\u0074\u0068\u0061\u0074\u0020s\u0070\u0061n\u0073\u0020\u006d\u0075\u006ct\u0069\u0070\u006c\u0065\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0061nd\u0020\u006c\u0069\u006b\u0065\u006cy\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u000a\u0065\u0078\u0061\u006d\u0070\u006c\u0065\u0073\u0020\u0061n\u0064\u0020\u0075s\u0061\u0067\u0065\u0020\u006ff\u0020\u0075\u0073\u0069\u006eg \u0079\u006f\u0075\u0072\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0046\u006f\u0072\u0020e\u0078\u0061\u006dp\u006c\u0065\u003a\u000a\u000a\u0043\u006f\u0062\u0072\u0061\u0020\u0069s\u0020\u0061\u0020\u0043L\u0049\u0020\u006c\u0069\u0062\u0072\u0061\u0072\u0079\u0020f\u006f\u0072\u0020\u0047\u006f\u0020t\u0068\u0061t\u0020\u0065\u006dp\u006f\u0077\u0065\u0072\u0073\u0020\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u002e\u000a\u0054\u0068\u0069\u0073\u0020\u0061\u0070p\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e \u0069\u0073 \u0061\u0020\u0074\u006f\u006f\u006c\u0020\u0074o\u0020g\u0065\u006e\u0065\u0072a\u0074\u0065 \u0074\u0068e\u0020\u006e\u0065\u0065\u0064\u0065\u0064\u0020\u0066\u0069\u006ce\u0073\u000a\u0074\u006f\u0020\u0071\u0075\u0069\u0063\u006b\u006cy\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0061\u0020\u0043\u006f\u0062\u0072\u0061\u0020\u0061\u0070pl\u0069ca\u0074\u0069on\u002e"};
type parametersConfig struct{

// PaperWidth sets the width of the paper.
PaperWidth _d .LengthFlag `mapstructure:"paper-width"`;

// PaperHeight is the height of the output paper.
PaperHeight _d .LengthFlag `mapstructure:"paper-height"`;

// PageSize is the page size string.
PageSize _d .PageSize `mapstructure:"page-size"`;

// Orientation defines if the output should be in a landscape format.
Orientation _d .Orientation `mapstructure:"orientation"`;

// MarginTop sets up the Top Margin for the output.
MarginTop _d .LengthFlag `mapstructure:"margin-top"`;

// MarginBottom sets up the Bottom Margin for the output.
MarginBottom _d .LengthFlag `mapstructure:"margin-bottom"`;

// MarginLeft sets up the Left Margin for the output.
MarginLeft _d .LengthFlag `mapstructure:"margin-left"`;

// MarginRight sets up the Right Margin for the output.
MarginRight _d .LengthFlag `mapstructure:"margin-right"`;};func _ff (){_eaa :=_ae .LogLevelInfo ;if _fd {_eaa =_ae .LogLevelDebug ;};if _ggb {_eaa =_ae .LogLevelTrace ;};_ae .Log =_ae .NewConsoleLogger (_eaa );};var (_fd ,_ggb bool ;);var _bb =&_a .Command {Use :"\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065",Short :"\u0047\u0065\u006e\u0065\u0072a\u0074\u0065\u0073\u0020\u0050\u0044F\u0020\u0062\u0061\u0073\u0065\u0064\u0020o\u006e\u0020\u0074h\u0065\u0020\u0070\u0072o\u0076\u0069\u0064\u0065\u0064\u0020H\u0054\u004d\u004c\u0020\u006f\u0072\u0020\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020w\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u0048\u0054\u004d\u004c\u0020\u0066\u0069\u006c\u0065\u0073\u002e",Long :"A\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069on\u0020\u0074\u0068\u0061\u0074\u0020s\u0070\u0061\u006e\u0073\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006c\u0069\u006e\u0065\u0073 \u0061\u006e\u0064\u0020\u006c\u0069\u006b\u0065l\u0079\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0065\u0078\u0061\u006d\u0070\u006c\u0065\u0073\u000a\u0061\u006e\u0064\u0020\u0075\u0073\u0061\u0067\u0065\u0020\u006f\u0066\u0020u\u0073\u0069\u006e\u0067\u0020\u0079o\u0075\u0072\u0020\u0063o\u006d\u006d\u0061\u006e\u0064\u002e\u0020\u0046\u006f\u0072\u0020e\u0078\u0061\u006d\u0070\u006c\u0065\u003a\u000a\u000a\u0043\u006f\u0062r\u0061\u0020\u0069\u0073\u0020\u0061\u0020\u0043\u004c\u0049\u0020\u006c\u0069\u0062\u0072\u0061r\u0079 \u0066\u006f\u0072\u0020\u0047\u006f\u0020\u0074\u0068\u0061\u0074\u0020\u0065\u006d\u0070\u006f\u0077\u0065\u0072\u0073\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u0073\u002e\u000a\u0054\u0068\u0069\u0073\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0069s\u0020\u0061\u0020\u0074\u006f\u006fl\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0074\u0068e\u0020n\u0065\u0065\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065s\u000a\u0074o\u0020\u0071\u0075\u0069\u0063\u006b\u006c\u0079\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0061\u0020C\u006fb\u0072\u0061\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074io\u006e\u002e",Run :_bc ,Args :_a .ExactArgs (2),ArgAliases :[]string {"\u0069\u006e\u0070u\u0074","\u006f\u0075\u0074\u0070\u0075\u0074\u002d\u0070\u0064\u0066"},Example :"\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0069\u006ep\u0075\u0074\u002e\u0068\u0074\u006d\u006c\u0020o\u0075t\u0070\u0075\u0074\u002e\u0070\u0064\u0066\u0020\u002d\u002d\u006f\u0072\u0069\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u0020\u0070\u006f\u0072\u0074\u0072\u0061\u0069\u0074"};
func init (){_a .OnInitialize (_fa );_gaf .PersistentFlags ().BoolVarP (&_fd ,"\u0064\u0065\u0062u\u0067","\u0064",false ,"\u0044e\u0066i\u006e\u0065\u0073\u0020\u0064e\u0062\u0075g\u0020\u006d\u006f\u0064\u0065");_gaf .PersistentFlags ().BoolVarP (&_ggb ,"\u0076e\u0072\u0062\u006f\u0073\u0065","\u0076",false ,"\u0056\u0065\u0072\u0062\u006f\u0073e\u0020\u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020o\u0066\u0020\u0074\u0068\u0065\u0020\u0063l\u0069\u0065\u006e\u0074");
_gaf .PersistentFlags ().StringVar (&_fe ,"\u0063\u006f\u006e\u0066\u0069\u0067","","\u0063\u006f\u006e\u0066\u0069\u0067\u0020\u0066i\u006c\u0065\u0020(d\u0065\u0066\u0061\u0075\u006c\u0074 \u0069\u0073\u0020\u0024\u0048\u004f\u004d\u0045\u002f\u002e\u0075\u006e\u0069\u0068\u0074m\u006c\u002d\u0073\u0072\u0063\u002e\u0079\u0061m\u006c\u0029");
_gaf .Flags ().BoolP ("\u0074\u006f\u0067\u0067\u006c\u0065","\u0074",false ,"\u0048\u0065\u006cp \u006d\u0065\u0073\u0073\u0061\u0067\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u006f\u0067\u0067\u006c\u0065");};var _fe string ;var (_c =generateConfig {};
_eafc =parametersConfig {PaperWidth :_d .LengthFlag {Length :_d .Inch (8.5).Millimeters ()},PaperHeight :_d .LengthFlag {Length :_d .Inch (11).Millimeters ()},Orientation :_d .Portrait ,MarginTop :_d .LengthFlag {Length :_d .Millimeter (10)},MarginBottom :_d .LengthFlag {Length :_d .Millimeter (10)},MarginLeft :_d .LengthFlag {Length :_d .Millimeter (10)},MarginRight :_d .LengthFlag {Length :_d .Millimeter (10)}};
);