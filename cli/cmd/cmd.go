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

package cmd ;import (_da "context";_dd "fmt";_cb "github.com/mitchellh/go-homedir";_a "github.com/spf13/cobra";_cee "github.com/spf13/viper";_dadd "github.com/unidoc/unihtml/client";_bc "github.com/unidoc/unihtml/content";_cf "github.com/unidoc/unihtml/sizes";_ce "github.com/unidoc/unipdf/v3/common";_c "os";_dad "path/filepath";_b "time";);func _gg (){if _dcb !=""{_cee .SetConfigFile (_dcb );}else {_beb ,_gb :=_cb .Dir ();if _gb !=nil {_dd .Println (_gb );_c .Exit (1);};_cee .AddConfigPath (_beb );_cee .SetConfigName ("\u002e\u0075\u006ei\u0068\u0074\u006d\u006c\u002d\u0073\u0072\u0063");};_cee .AutomaticEnv ();if _gbb :=_cee .ReadInConfig ();_gbb ==nil {_dd .Println ("\u0055s\u0069n\u0067\u0020\u0063\u006f\u006ef\u0069\u0067 \u0066\u0069\u006c\u0065\u003a",_cee .ConfigFileUsed ());};};var _bf =&_a .Command {Use :"\u0075n\u0069\u0068\u0074\u006d\u006c",Short :"\u0041\u0020\u0062\u0072\u0069\u0065\u0066\u0020\u0064\u0065\u0073\u0063\u0072i\u0070\u0074\u0069\u006f\u006e\u0020o\u0066\u0020\u0079\u006f\u0075\u0072\u0020\u0061\u0070\u0070\u006c\u0069\u0063a\u0074\u0069\u006f\u006e",Long :"\u0041\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0064\u0065\u0073c\u0072\u0069\u0070\u0074i\u006f\u006e\u0020\u0074\u0068\u0061\u0074\u0020s\u0070\u0061n\u0073\u0020\u006d\u0075\u006ct\u0069\u0070\u006c\u0065\u0020\u006c\u0069\u006e\u0065\u0073\u0020\u0061nd\u0020\u006c\u0069\u006b\u0065\u006cy\u0020\u0063\u006f\u006e\u0074\u0061\u0069\u006e\u0073\u000a\u0065\u0078\u0061\u006d\u0070\u006c\u0065\u0073\u0020\u0061n\u0064\u0020\u0075s\u0061\u0067\u0065\u0020\u006ff\u0020\u0075\u0073\u0069\u006eg \u0079\u006f\u0075\u0072\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006fn\u002e\u0020\u0046\u006f\u0072\u0020e\u0078\u0061\u006dp\u006c\u0065\u003a\u000a\u000a\u0043\u006f\u0062\u0072\u0061\u0020\u0069s\u0020\u0061\u0020\u0043L\u0049\u0020\u006c\u0069\u0062\u0072\u0061\u0072\u0079\u0020f\u006f\u0072\u0020\u0047\u006f\u0020t\u0068\u0061t\u0020\u0065\u006dp\u006f\u0077\u0065\u0072\u0073\u0020\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0073\u002e\u000a\u0054\u0068\u0069\u0073\u0020\u0061\u0070p\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e \u0069\u0073 \u0061\u0020\u0074\u006f\u006f\u006c\u0020\u0074o\u0020g\u0065\u006e\u0065\u0072a\u0074\u0065 \u0074\u0068e\u0020\u006e\u0065\u0065\u0064\u0065\u0064\u0020\u0066\u0069\u006ce\u0073\u000a\u0074\u006f\u0020\u0071\u0075\u0069\u0063\u006b\u006cy\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0061\u0020\u0043\u006f\u0062\u0072\u0061\u0020\u0061\u0070pl\u0069ca\u0074\u0069on\u002e"};var _dcb string ;type parametersConfig struct{

// PaperWidth sets the width of the paper.
PaperWidth _cf .LengthFlag `mapstructure:"paper-width"`;

// PaperHeight is the height of the output paper.
PaperHeight _cf .LengthFlag `mapstructure:"paper-height"`;

// PageSize is the page size string.
PageSize _cf .PageSize `mapstructure:"page-size"`;

// Orientation defines if the output should be in a landscape format.
Orientation _cf .Orientation `mapstructure:"orientation"`;

// MarginTop sets up the Top Margin for the output.
MarginTop _cf .LengthFlag `mapstructure:"margin-top"`;

// MarginBottom sets up the Bottom Margin for the output.
MarginBottom _cf .LengthFlag `mapstructure:"margin-bottom"`;

// MarginLeft sets up the Left Margin for the output.
MarginLeft _cf .LengthFlag `mapstructure:"margin-left"`;

// MarginRight sets up the Right Margin for the output.
MarginRight _cf .LengthFlag `mapstructure:"margin-right"`;};func init (){_bf .AddCommand (_be );_be .Flags ().IntP ("\u0070\u006f\u0072\u0074","\u0070",8080,"\u0050\u006f\u0072\u0074\u0020\u006f\u0066\u0020\u0074\u0068\u0065 \u0075\u006e\u0069\u0068\u0074\u006d\u006c\u0020\u0073\u0065r\u0076\u0065\u0072");_be .Flags ().String ("\u0068\u006f\u0073\u0074","\u006co\u0063\u0061\u006c\u0068\u006f\u0073t","\u0048\u006f\u0073t\u0020\u006e\u0061\u006de\u0020\u006f\u0066\u0020\u0074\u0068\u0065 \u0075\u006e\u0069\u0068\u0074\u006d\u006c\u0020\u0073\u0065\u0072\u0076\u0065\u0072");_be .Flags ().BoolP ("\u0068\u0074\u0074p\u0073","\u0073",false ,"\u0050\u0072o\u0074\u006f\u0063\u006fl\u0020\u0075s\u0065\u0064\u0020\u0069\u006e\u0020\u0073\u0065r\u0076\u0065\u0072\u0020\u0063\u006f\u006d\u006d\u0075\u006e\u0069\u0063a\u0074\u0069\u006f\u006e");_be .Flags ().StringP ("\u0070\u0072\u0065\u0066\u0069\u0078","\u0078","","\u0050u\u0062\u006ci\u0063\u0020\u0061\u0070i\u0020\u0070\u0072e\u0066\u0069\u0078\u0020\u0075\u0073\u0065\u0064\u0020by\u0020\u0074\u0068e\u0020\u0075n\u0069\u0068\u0074\u006d\u006c\u0020s\u0065\u0072v\u0065\u0072");_be .Flags ().Var (&_f .PaperWidth ,"p\u0061\u0070\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068","\u0073\u0065\u0074s \u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0070\u0065\u0072\u002d\u0077\u0069\u0064\u0074\u0068");_be .Flags ().Var (&_f .PaperHeight ,"\u0070\u0061\u0070e\u0072\u002d\u0068\u0065\u0069\u0067\u0068\u0074","\u0073e\u0074\u0073\u0020\u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070a\u0070\u0065\u0072\u002d\u0068\u0065\u0069\u0067\u0068\u0074");_be .Flags ().Var (&_f .PageSize ,"\u0070\u0061\u0070\u0065\u0072\u002d\u0073\u0069\u007a\u0065","s\u0065\u0074\u0073\u0020up\u0020t\u0068\u0065\u0020\u0070\u0061g\u0065\u0020\u0073\u0069\u007a\u0065");_be .Flags ().Var (&_f .Orientation ,"o\u0072\u0069\u0065\u006e\u0074\u0061\u0074\u0069\u006f\u006e","\u0073\u0065\u0074\u0073 \u0075\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0067e\u0020o\u0072\u0069\u0065\u006e\u0074\u0061\u0074i\u006f\u006e");_be .Flags ().Var (&_f .MarginTop ,"\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074\u006f\u0070","\u0073\u0065\u0074\u0073 u\u0070\u0020\u0074\u0068\u0065\u0020\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0074o\u0070");_be .Flags ().Var (&_f .MarginBottom ,"\u006d\u0061\u0072\u0067\u0069\u006e\u002d\u0062\u006f\u0074\u0074\u006f\u006d","\u0073e\u0074\u0073\u0020\u0075p\u0020\u0074\u0068\u0065\u0020m\u0061r\u0067i\u006e\u002d\u0062\u006f\u0074\u0074\u006fm");_be .Flags ().Var (&_f .MarginRight ,"\u006d\u0061\u0072g\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074","\u0073e\u0074\u0073\u0020\u0075\u0070\u0020\u0074\u0068\u0065\u0020\u006da\u0072\u0067\u0069\u006e\u002d\u0072\u0069\u0067\u0068\u0074");_be .Flags ().Var (&_f .MarginLeft ,"m\u0061\u0072\u0067\u0069\u006e\u002d\u006c\u0065\u0066\u0074","\u0073\u0065\u0074\u0073 u\u0070\u0020\u0074\u0068\u0065\u0020\u0070\u0061\u0070\u0065\u0072\u002d\u006c\u0065f\u0074");};var (_dae =generateConfig {};_f =parametersConfig {PaperWidth :_cf .LengthFlag {Length :_cf .Inch (8.5).Millimeters ()},PaperHeight :_cf .LengthFlag {Length :_cf .Inch (11).Millimeters ()},Orientation :_cf .Portrait ,MarginTop :_cf .LengthFlag {Length :_cf .Millimeter (10)},MarginBottom :_cf .LengthFlag {Length :_cf .Millimeter (10)},MarginLeft :_cf .LengthFlag {Length :_cf .Millimeter (10)},MarginRight :_cf .LengthFlag {Length :_cf .Millimeter (10)}};);var _be =&_a .Command {Use :"\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065",Short :"\u0047\u0065\u006e\u0065\u0072a\u0074\u0065\u0073\u0020\u0050\u0044F\u0020\u0062\u0061\u0073\u0065\u0064\u0020o\u006e\u0020\u0074h\u0065\u0020\u0070\u0072o\u0076\u0069\u0064\u0065\u0064\u0020H\u0054\u004d\u004c\u0020\u006f\u0072\u0020\u0064\u0069\u0072\u0065\u0063\u0074\u006f\u0072\u0079\u0020w\u0069\u0074\u0068\u0020\u0074\u0068\u0065\u0020\u0048\u0054\u004d\u004c\u0020\u0066\u0069\u006c\u0065\u0073\u002e",Long :"A\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069on\u0020\u0074\u0068\u0061\u0074\u0020s\u0070\u0061\u006e\u0073\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020\u006c\u0069\u006e\u0065\u0073 \u0061\u006e\u0064\u0020\u006c\u0069\u006b\u0065l\u0079\u0020\u0063o\u006e\u0074\u0061\u0069\u006e\u0073\u0020\u0065\u0078\u0061\u006d\u0070\u006c\u0065\u0073\u000a\u0061\u006e\u0064\u0020\u0075\u0073\u0061\u0067\u0065\u0020\u006f\u0066\u0020u\u0073\u0069\u006e\u0067\u0020\u0079o\u0075\u0072\u0020\u0063o\u006d\u006d\u0061\u006e\u0064\u002e\u0020\u0046\u006f\u0072\u0020e\u0078\u0061\u006d\u0070\u006c\u0065\u003a\u000a\u000a\u0043\u006f\u0062r\u0061\u0020\u0069\u0073\u0020\u0061\u0020\u0043\u004c\u0049\u0020\u006c\u0069\u0062\u0072\u0061r\u0079 \u0066\u006f\u0072\u0020\u0047\u006f\u0020\u0074\u0068\u0061\u0074\u0020\u0065\u006d\u0070\u006f\u0077\u0065\u0072\u0073\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u0073\u002e\u000a\u0054\u0068\u0069\u0073\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0069s\u0020\u0061\u0020\u0074\u006f\u006fl\u0020\u0074\u006f\u0020\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0074\u0068e\u0020n\u0065\u0065\u0064\u0065\u0064\u0020\u0066\u0069\u006c\u0065s\u000a\u0074o\u0020\u0071\u0075\u0069\u0063\u006b\u006c\u0079\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020\u0061\u0020C\u006fb\u0072\u0061\u0020\u0061\u0070\u0070\u006c\u0069\u0063\u0061\u0074io\u006e\u002e",Run :_ced ,Args :_a .ExactArgs (2),ArgAliases :[]string {"\u0069\u006e\u0070u\u0074","\u006f\u0075\u0074\u0070\u0075\u0074\u002d\u0070\u0064\u0066"},Example :"\u0067\u0065\u006e\u0065\u0072\u0061\u0074\u0065\u0020\u0069\u006ep\u0075\u0074\u002e\u0068\u0074\u006d\u006c\u0020o\u0075t\u0070\u0075\u0074\u002e\u0070\u0064\u0066\u0020\u002d\u002d\u006f\u0072\u0069\u0065\u006e\u0074\u0061\u0074i\u006f\u006e\u0020\u0070\u006f\u0072\u0074\u0072\u0061\u0069\u0074"};

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute (){if _fd :=_bf .Execute ();_fd !=nil {_dd .Println (_fd );_c .Exit (1);};};func _de (){_cg :=_ce .LogLevelInfo ;if _ef {_cg =_ce .LogLevelDebug ;};if _ac {_cg =_ce .LogLevelTrace ;};_ce .Log =_ce .NewConsoleLogger (_cg );};type generateConfig struct{Port int `mapstructure:"port"`;Host string `mapstructure:"host"`;Https bool `mapstructure:"https"`;Prefix string `mapstructure:"prefix"`;};func init (){_a .OnInitialize (_gg );_bf .PersistentFlags ().BoolVarP (&_ef ,"\u0064\u0065\u0062u\u0067","\u0064",false ,"\u0044e\u0066i\u006e\u0065\u0073\u0020\u0064e\u0062\u0075g\u0020\u006d\u006f\u0064\u0065");_bf .PersistentFlags ().BoolVarP (&_ac ,"\u0076e\u0072\u0062\u006f\u0073\u0065","\u0076",false ,"\u0056\u0065\u0072\u0062\u006f\u0073e\u0020\u0069\u006e\u0066\u006f\u0072\u006d\u0061\u0074\u0069\u006f\u006e\u0020o\u0066\u0020\u0074\u0068\u0065\u0020\u0063l\u0069\u0065\u006e\u0074");_bf .PersistentFlags ().StringVar (&_dcb ,"\u0063\u006f\u006e\u0066\u0069\u0067","","\u0063\u006f\u006e\u0066\u0069\u0067\u0020\u0066i\u006c\u0065\u0020(d\u0065\u0066\u0061\u0075\u006c\u0074 \u0069\u0073\u0020\u0024\u0048\u004f\u004d\u0045\u002f\u002e\u0075\u006e\u0069\u0068\u0074m\u006c\u002d\u0073\u0072\u0063\u002e\u0079\u0061m\u006c\u0029");_bf .Flags ().BoolP ("\u0074\u006f\u0067\u0067\u006c\u0065","\u0074",false ,"\u0048\u0065\u006cp \u006d\u0065\u0073\u0073\u0061\u0067\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u006f\u0067\u0067\u006c\u0065");};func _ced (cmd *_a .Command ,_dc []string ){_g :=_b .Now ();if _ad :=_cee .BindPFlags (cmd .Flags ());_ad !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ad );_c .Exit (1);};if _daf :=_cee .Unmarshal (&_dae );_daf !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_daf );_c .Exit (1);};_de ();_aa :=_b .Now ();_gf ,_ga :=_c .Stat (_dc [0]);if _ga !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ga );_c .Exit (1);};if !_gf .IsDir (){if _dad .Ext (_gf .Name ())!="\u002e\u0068\u0074m\u006c"{_dd .Printf ("\u0045\u0072r\u003a\u0020\u0043\u0075\u0072\u0072\u0065\u006e\u0074\u006c\u0079\u0020\u006f\u006e\u006c\u0079\u0020\u0048\u0054M\u004c\u0020\u0066\u0069\u006c\u0065s\u0020\u0069\u006e\u0070\u0075\u0074\u0053\u0074\u0061\u0074\u0020\u0061\u0072\u0065 \u0073\u0075p\u0070\u006f\u0072\u0074e\u0064\u002e\u0020\u0049\u006ep\u0075\u0074\u003a\u0020\u0025\u0073\u000a",_dc [0]);_c .Exit (1);};};_df ,_ga :=_c .OpenFile (_dc [1],_c .O_CREATE |_c .O_WRONLY |_c .O_TRUNC ,0700);if _ga !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ga );_c .Exit (1);};defer _df .Close ();_gfa :=_dadd .New (_dadd .Options {HTTPS :_dae .Https ,Hostname :_dae .Host ,Port :_dae .Port ,Prefix :_dae .Prefix });_adf ,_gab :=_da .WithTimeout (_da .Background (),_b .Second *10);defer _gab ();_aa =_b .Now ();var _ba _bc .Content ;if _gf .IsDir (){_ba ,_ga =_bc .NewZipDirectory (_dc [0]);}else {_ba ,_ga =_bc .NewHTMLFile (_dc [0]);};if _ga !=nil {_dd .Printf ("\u0045r\u0072\u003a\u0020\u0025\u0076",_ga );_c .Exit (1);};_bcc ,_ga :=_dadd .BuildHTMLQuery ().PaperWidth (_f .PaperWidth .Length ).PaperHeight (_f .PaperHeight .Length ).PageSize (_f .PageSize ).MarginTop (_f .MarginTop .Length ).MarginBottom (_f .MarginBottom .Length ).MarginLeft (_f .MarginLeft .Length ).MarginRight (_f .MarginRight .Length ).Orientation (_f .Orientation ).SetContent (_ba ).Query ();if _ga !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ga );_c .Exit (1);};_ee ,_ga :=_gfa .ConvertHTML (_adf ,_bcc );if _ga !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ga );_c .Exit (1);};_ce .Log .Trace ("\u0045\u0078\u0065cu\u0074\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0065r\u0061t\u0065 \u0071u\u0065\u0072\u0079\u0020\u0074\u0061\u006b\u0065\u006e\u003a\u0020\u0025\u0073",_b .Since (_aa ));_aa =_b .Now ();_ ,_ga =_df .Write (_ee .Data );if _ga !=nil {_dd .Printf ("\u0045\u0072\u0072\u003a\u0020\u0025\u0076\u000a",_ga );_c .Exit (1);};_ce .Log .Trace ("\u0057\u0072\u0069\u0074in\u0067\u0020\u0066\u0069\u006c\u0065\u0020\u0074\u0061\u006b\u0065\u006e\u003a\u0020%\u0073",_b .Since (_aa ));_dd .Printf ("\u0047\u0065n\u0065\u0072\u0061\u0074\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0073\u0075\u0063\u0063\u0065\u0073\u0073\u0020\u0069\u006e %\u0073\u000a",_b .Since (_g ));};var (_ef ,_ac bool ;);