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

package decoder ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_c "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_fe "github.com/unidoc/unipdf/v3/internal/jbig2/document";_b "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_a "image";
);func (_d *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _d .decodePage (pageNumber )};type Decoder struct{_e *_f .Reader ;_fb *_fe .Document ;_bg int ;_cf Parameters ;};func Decode (input []byte ,parameters Parameters ,globals *_fe .Globals )(*Decoder ,error ){_cc :=_f .NewReader (input );
_bf ,_cg :=_fe .DecodeDocument (_cc ,globals );if _cg !=nil {return nil ,_cg ;};return &Decoder {_e :_cc ,_fb :_bf ,_cf :parameters },nil ;};func (_bc *Decoder )decodePage (_ed int )([]byte ,error ){const _ae ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";
if _ed < 0{return nil ,_b .Errorf (_ae ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ed );};if _ed > int (_bc ._fb .NumberOfPages ){return nil ,_b .Errorf (_ae ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ed );
};_gf ,_dd :=_bc ._fb .GetPage (_ed );if _dd !=nil {return nil ,_b .Wrap (_dd ,_ae ,"");};_ab ,_dd :=_gf .GetBitmap ();if _dd !=nil {return nil ,_b .Wrap (_dd ,_ae ,"");};_ab .InverseData ();if !_bc ._cf .UnpaddedData {return _ab .Data ,nil ;};return _ab .GetUnpaddedData ();
};func (_de *Decoder )PageNumber ()(int ,error ){const _ca ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _de ._fb ==nil {return 0,_b .Error (_ca ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");
};return int (_de ._fb .NumberOfPages ),nil ;};func (_ge *Decoder )DecodeNextPage ()([]byte ,error ){_ge ._bg ++;_ga :=_ge ._bg ;return _ge .decodePage (_ga );};func (_dc *Decoder )decodePageImage (_eg int )(_a .Image ,error ){const _ba ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";
if _eg < 0{return nil ,_b .Errorf (_ba ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_eg );};if _eg > int (_dc ._fb .NumberOfPages ){return nil ,_b .Errorf (_ba ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_eg );
};_cd ,_ag :=_dc ._fb .GetPage (_eg );if _ag !=nil {return nil ,_b .Wrap (_ag ,_ba ,"");};_ddb ,_ag :=_cd .GetBitmap ();if _ag !=nil {return nil ,_b .Wrap (_ag ,_ba ,"");};_ddb .InverseData ();return _ddb .ToImage (),nil ;};type Parameters struct{UnpaddedData bool ;
Color _c .Color ;};func (_cb *Decoder )DecodePageImage (pageNumber int )(_a .Image ,error ){const _ee ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";_gc ,_ad :=_cb .decodePageImage (pageNumber );
if _ad !=nil {return nil ,_b .Wrap (_ad ,_ee ,"");};return _gc ,nil ;};