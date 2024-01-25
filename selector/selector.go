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

package selector ;import _g "errors";

// ByType is a 'By' selector type enumerator.
type ByType uint ;

// Validate checks validity of the ByType.
func (_cg ByType )Validate ()error {if _cg >=ByID &&_cg <=BySearch {return nil ;};return _g .New ("\u0069\u006e\u0076\u0061li\u0064\u0020\u0062\u0079\u0020\u0073\u0065\u006c\u0065\u0063\u0074\u006f\u0072");};const (ByUndefined ByType =iota ;ByID ;ByQueryAll ;
ByQuery ;ByNodeID ;ByJSPath ;BySearch ;);