package mosaic

import (
	"image"
	"image/png"
	"os"
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	raw, err := loadImage("./fixtures/charm-wish.png")
	if err != nil {
		os.Exit(1)
	}

	// Create options
	var mosaic = New().Width(80).Height(40)

	result := mosaic.Render(raw)

	// TODO: Use a smaller image, it is almost impossible seeing the diff by now
	expected := `[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;189;156;174;48;2;0;0;0m▄[m[38;2;233;192;214;48;2;0;0;0m█[m[38;2;254;210;234;48;2;0;0;0m█[m[38;2;254;209;233;48;2;0;0;0m█[m[38;2;255;209;233;48;2;56;47;52m▄[m[38;2;0;0;0;48;2;21;17;20m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;23;19;21m [m[38;2;253;208;228;48;2;0;0;0m█[m[38;2;253;208;226;48;2;0;0;0m█[m[38;2;253;206;225;48;2;0;0;0m█[m[38;2;253;205;226;48;2;0;0;0m█[m[38;2;254;204;227;48;2;0;0;0m█[m[38;2;141;114;128;48;2;175;140;158m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;28;23;26m [m[38;2;255;210;234;48;2;0;0;0m▄[m[38;2;255;210;235;48;2;0;0;0m▄[m[38;2;118;97;108;48;2;0;0;0m▄[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;252;204;223;48;2;94;75;82m▀[m[38;2;251;201;216;48;2;0;0;0m█[m[38;2;251;198;214;48;2;0;0;0m█[m[38;2;252;196;215;48;2;0;0;0m█[m[38;2;252;193;218;48;2;0;0;0m█[m[38;2;128;100;113;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;140;112;124;48;2;63;51;57m▄[m[38;2;253;206;227;48;2;0;0;0m█[m[38;2;254;207;230;48;2;0;0;0m█[m[38;2;254;208;232;48;2;0;0;0m█[m[38;2;254;208;233;48;2;0;0;0m█[m[38;2;110;90;101;48;2;175;143;161m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;232;226;154;48;2;0;0;0m▄[m[38;2;254;249;169;48;2;68;66;45m▄[m[38;2;254;247;168;48;2;0;0;0m▄[m[38;2;84;81;55;48;2;1;1;1m▄[m[38;2;0;0;0;48;2;51;38;43m [m[38;2;0;0;0;48;2;42;32;36m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;115;92;105;48;2;0;0;0m▄[m[38;2;255;206;232;48;2;0;0;0m▄[m[38;2;255;207;233;48;2;0;0;0m▄[m[38;2;243;197;222;48;2;0;0;0m▄[m[38;2;0;0;0;48;2;32;26;29m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;187;146;162;48;2;127;94;107m▀[m[38;2;251;196;217;48;2;0;0;0m█[m[38;2;252;198;220;48;2;0;0;0m█[m[38;2;252;199;224;48;2;0;0;0m█[m[38;2;254;201;227;48;2;0;0;0m█[m[38;2;163;132;148;48;2;83;66;75m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;3;3;2m [m[38;2;253;249;170;48;2;126;124;85m▄[m[38;2;253;249;170;48;2;0;0;0m█[m[38;2;253;249;169;48;2;0;0;0m█[m[38;2;253;248;169;48;2;0;0;0m█[m[38;2;253;245;167;48;2;0;0;0m█[m[38;2;137;131;89;48;2;1;1;0m▄[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;37;29;32m [m[38;2;254;203;223;48;2;93;74;83m▄[m[38;2;254;204;225;48;2;0;0;0m█[m[38;2;254;204;226;48;2;0;0;0m█[m[38;2;254;205;227;48;2;0;0;0m█[m[38;2;254;205;229;48;2;0;0;0m█[m[38;2;254;205;230;48;2;0;0;0m█[m[38;2;254;204;230;48;2;0;0;0m█[m[38;2;255;203;230;48;2;0;0;0m█[m[38;2;254;200;228;48;2;78;63;71m▄[m[38;2;110;85;98;48;2;0;0;0m▄[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;164;118;137;48;2;0;0;0m▀[m[38;2;251;185;211;48;2;0;0;0m▀[m[38;2;252;188;217;48;2;0;0;0m▀[m[38;2;99;76;87;48;2;0;0;0m▀[m[38;2;92;72;80;48;2;0;0;0m▄[m[38;2;254;203;227;48;2;0;0;0m▄[m[38;2;255;203;231;48;2;0;0;0m▄[m[38;2;132;106;120;48;2;0;0;0m▄[m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;198;159;180;48;2;0;0;0m▄[m[38;2;254;207;224;48;2;0;0;0m▄[m[38;2;250;249;172;48;2;136;136;92m▄[m[38;2;250;250;171;48;2;0;0;0m█[m[38;2;250;250;171;48;2;0;0;0m█[m[38;2;250;249;171;48;2;0;0;0m█[m[38;2;250;249;171;48;2;0;0;0m█[m[38;2;250;248;170;48;2;0;0;0m█[m[38;2;250;245;167;48;2;0;0;0m█[m[38;2;240;219;175;48;2;16;15;10m▄[m[38;2;245;192;207;48;2;0;0;0m█[m[38;2;251;196;208;48;2;0;0;0m█[m[38;2;251;197;209;48;2;0;0;0m█[m[38;2;252;199;212;48;2;0;0;0m█[m[38;2;253;200;216;48;2;0;0;0m█[m[38;2;252;201;219;48;2;0;0;0m█[m[38;2;253;201;223;48;2;0;0;0m█[m[38;2;253;201;225;48;2;0;0;0m█[m[38;2;253;200;225;48;2;0;0;0m█[m[38;2;254;199;226;48;2;0;0;0m█[m[38;2;254;196;225;48;2;0;0;0m█[m[38;2;142;108;124;48;2;18;14;16m▄[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;162;125;138;48;2;119;88;99m▀[m[38;2;252;197;220;48;2;0;0;0m█[m[38;2;254;201;227;48;2;0;0;0m█[m[38;2;251;200;227;48;2;127;100;114m▀[m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;29;23;26m [m[38;2;254;202;229;48;2;105;84;96m▄[m[38;2;254;202;227;48;2;0;0;0m█[m[38;2;252;210;216;48;2;0;0;0m█[m[38;2;247;241;183;48;2;0;0;0m█[m[38;2;246;250;173;48;2;0;0;0m█[m[38;2;246;250;173;48;2;0;0;0m█[m[38;2;246;250;173;48;2;0;0;0m█[m[38;2;246;250;173;48;2;0;0;0m█[m[38;2;246;250;173;48;2;0;0;0m█[m[38;2;247;250;172;48;2;0;0;0m█[m[38;2;247;249;172;48;2;0;0;0m█[m[38;2;247;247;170;48;2;0;0;0m█[m[38;2;247;226;177;48;2;0;0;0m█[m[38;2;244;184;184;48;2;0;0;0m█[m[38;2;245;185;188;48;2;0;0;0m█[m[38;2;246;188;193;48;2;0;0;0m█[m[38;2;247;191;201;48;2;0;0;0m█[m[38;2;249;193;207;48;2;0;0;0m█[m[38;2;250;194;212;48;2;0;0;0m█[m[38;2;251;195;215;48;2;0;0;0m█[m[38;2;252;195;219;48;2;0;0;0m█[m[38;2;253;194;220;48;2;0;0;0m█[m[38;2;253;193;222;48;2;0;0;0m█[m[38;2;254;191;222;48;2;0;0;0m█[m[38;2;118;88;103;48;2;0;0;0m▄[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;1;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;6;4;5m [m[38;2;254;199;226;48;2;199;158;179m▄[m[38;2;252;199;222;48;2;0;0;0m█[m[38;2;246;224;198;48;2;0;0;0m█[m[38;2;241;250;175;48;2;0;0;0m█[m[38;2;241;251;175;48;2;0;0;0m█[m[38;2;241;250;175;48;2;0;0;0m█[m[38;2;241;250;175;48;2;0;0;0m█[m[38;2;241;250;175;48;2;0;0;0m█[m[38;2;241;250;175;48;2;0;0;0m█[m[38;2;242;250;175;48;2;0;0;0m█[m[38;2;242;250;175;48;2;0;0;0m█[m[38;2;242;250;174;48;2;0;0;0m█[m[38;2;242;250;175;48;2;0;0;0m█[m[38;2;243;249;174;48;2;0;0;0m█[m[38;2;243;244;173;48;2;0;0;0m█[m[38;2;239;212;173;48;2;0;0;0m█[m[38;2;238;189;175;48;2;0;0;0m█[m[38;2;239;178;182;48;2;0;0;0m█[m[38;2;242;181;190;48;2;0;0;0m█[m[38;2;243;183;196;48;2;0;0;0m█[m[38;2;246;185;202;48;2;0;0;0m█[m[38;2;248;187;207;48;2;0;0;0m█[m[38;2;249;187;211;48;2;0;0;0m█[m[38;2;251;189;214;48;2;0;0;0m█[m[38;2;252;189;216;48;2;0;0;0m█[m[38;2;252;189;215;48;2;0;0;0m█[m[38;2;252;190;217;48;2;0;0;0m█[m[38;2;253;192;218;48;2;0;0;0m█[m[38;2;253;193;221;48;2;0;0;0m█[m[38;2;254;193;222;48;2;0;0;0m█[m[38;2;254;191;222;48;2;0;0;0m█[m[38;2;248;186;217;48;2;0;0;0m█[m[38;2;254;191;223;48;2;14;11;12m▄[m[38;2;0;0;0;48;2;29;22;25m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;24;26;18m [m[38;2;234;251;178;48;2;51;40;46m▄[m[38;2;241;230;196;48;2;0;0;0m█[m[38;2;235;251;178;48;2;0;0;0m█[m[38;2;235;251;178;48;2;0;0;0m█[m[38;2;235;251;178;48;2;0;0;0m█[m[38;2;235;251;178;48;2;0;0;0m█[m[38;2;235;251;178;48;2;0;0;0m█[m[38;2;235;251;177;48;2;0;0;0m█[m[38;2;236;251;177;48;2;0;0;0m█[m[38;2;236;251;178;48;2;0;0;0m█[m[38;2;236;251;177;48;2;0;0;0m█[m[38;2;236;250;177;48;2;0;0;0m█[m[38;2;236;250;177;48;2;0;0;0m█[m[38;2;236;251;177;48;2;0;0;0m█[m[38;2;237;250;177;48;2;0;0;0m█[m[38;2;237;251;177;48;2;0;0;0m█[m[38;2;237;250;177;48;2;0;0;0m█[m[38;2;237;250;177;48;2;0;0;0m█[m[38;2;237;250;177;48;2;0;0;0m█[m[38;2;238;250;176;48;2;0;0;0m█[m[38;2;238;240;177;48;2;0;0;0m█[m[38;2;238;221;182;48;2;0;0;0m█[m[38;2;240;214;187;48;2;0;0;0m█[m[38;2;241;215;189;48;2;0;0;0m█[m[38;2;242;214;190;48;2;0;0;0m█[m[38;2;246;181;200;48;2;0;0;0m█[m[38;2;247;183;204;48;2;0;0;0m█[m[38;2;249;186;208;48;2;0;0;0m█[m[38;2;251;190;214;48;2;0;0;0m█[m[38;2;253;191;217;48;2;0;0;0m█[m[38;2;253;192;221;48;2;0;0;0m█[m[38;2;254;192;223;48;2;0;0;0m█[m[38;2;254;192;224;48;2;0;0;0m█[m[38;2;254;191;224;48;2;0;0;0m█[m[38;2;253;191;223;48;2;0;0;0m█[m[38;2;124;94;110;48;2;0;0;0m▄[m
[38;2;225;251;181;48;2;0;0;0m▄[m[38;2;210;231;166;48;2;0;0;0m█[m[38;2;228;251;181;48;2;0;0;0m█[m[38;2;228;251;181;48;2;0;0;0m█[m[38;2;228;251;181;48;2;0;0;0m█[m[38;2;229;251;180;48;2;0;0;0m█[m[38;2;229;251;180;48;2;0;0;0m█[m[38;2;229;251;181;48;2;0;0;0m█[m[38;2;229;251;180;48;2;0;0;0m█[m[38;2;230;251;180;48;2;0;0;0m█[m[38;2;229;251;180;48;2;0;0;0m█[m[38;2;229;251;180;48;2;0;0;0m█[m[38;2;230;251;180;48;2;0;0;0m█[m[38;2;230;251;180;48;2;0;0;0m█[m[38;2;230;251;180;48;2;0;0;0m█[m[38;2;230;251;180;48;2;0;0;0m█[m[38;2;231;251;180;48;2;0;0;0m█[m[38;2;231;251;179;48;2;0;0;0m█[m[38;2;231;251;179;48;2;0;0;0m█[m[38;2;231;251;180;48;2;0;0;0m█[m[38;2;231;251;179;48;2;0;0;0m█[m[38;2;232;251;179;48;2;0;0;0m█[m[38;2;231;251;179;48;2;0;0;0m█[m[38;2;232;251;179;48;2;0;0;0m█[m[38;2;232;251;179;48;2;0;0;0m█[m[38;2;232;251;179;48;2;0;0;0m█[m[38;2;233;251;179;48;2;0;0;0m█[m[38;2;233;251;179;48;2;0;0;0m█[m[38;2;234;249;178;48;2;0;0;0m█[m[38;2;238;202;185;48;2;0;0;0m█[m[38;2;242;177;196;48;2;0;0;0m█[m[38;2;245;181;203;48;2;0;0;0m█[m[38;2;248;184;208;48;2;0;0;0m█[m[38;2;250;187;214;48;2;0;0;0m█[m[38;2;252;189;218;48;2;0;0;0m█[m[38;2;254;189;221;48;2;0;0;0m█[m[38;2;254;190;222;48;2;0;0;0m█[m[38;2;255;190;223;48;2;0;0;0m█[m[38;2;254;190;223;48;2;0;0;0m█[m[38;2;235;175;205;48;2;167;125;147m▄[m
[38;2;175;198;144;48;2;50;57;42m▀[m[38;2;221;252;184;48;2;0;0;0m█[m[38;2;221;252;184;48;2;0;0;0m█[m[38;2;221;251;183;48;2;0;0;0m█[m[38;2;222;251;183;48;2;0;0;0m█[m[38;2;203;230;169;48;2;0;0;0m█[m[38;2;202;226;166;48;2;70;74;65m▀[m[38;2;222;251;183;48;2;0;0;0m█[m[38;2;222;251;183;48;2;0;0;0m█[m[38;2;222;251;183;48;2;0;0;0m█[m[38;2;222;252;183;48;2;0;0;0m█[m[38;2;223;251;183;48;2;0;0;0m█[m[38;2;223;251;183;48;2;0;0;0m█[m[38;2;223;252;182;48;2;0;0;0m█[m[38;2;0;0;0;48;2;53;54;51m [m[38;2;224;251;182;48;2;0;0;0m█[m[38;2;224;251;182;48;2;0;0;0m█[m[38;2;224;251;182;48;2;0;0;0m█[m[38;2;224;251;182;48;2;0;0;0m█[m[38;2;225;251;182;48;2;0;0;0m█[m[38;2;225;251;182;48;2;0;0;0m█[m[38;2;225;251;182;48;2;0;0;0m█[m[38;2;226;251;182;48;2;0;0;0m█[m[38;2;226;251;181;48;2;0;0;0m█[m[38;2;226;251;182;48;2;0;0;0m█[m[38;2;226;251;181;48;2;0;0;0m█[m[38;2;226;250;180;48;2;0;0;0m█[m[38;2;227;244;179;48;2;0;0;0m█[m[38;2;235;190;184;48;2;0;0;0m█[m[38;2;238;171;189;48;2;0;0;0m█[m[38;2;239;174;193;48;2;0;0;0m█[m[38;2;243;178;199;48;2;0;0;0m█[m[38;2;246;181;206;48;2;0;0;0m█[m[38;2;248;183;211;48;2;0;0;0m█[m[38;2;251;185;215;48;2;0;0;0m█[m[38;2;253;186;218;48;2;0;0;0m█[m[38;2;253;187;220;48;2;0;0;0m█[m[38;2;254;187;221;48;2;0;0;0m█[m[38;2;255;187;221;48;2;0;0;0m█[m[38;2;171;126;149;48;2;124;91;108m▀[m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;112;131;96;48;2;0;0;0m▀[m[38;2;212;251;185;48;2;0;0;0m█[m[38;2;213;252;186;48;2;0;0;0m█[m[38;2;214;252;186;48;2;0;0;0m█[m[38;2;176;204;153;48;2;163;191;145m▀[m[38;2;85;93;79;48;2;79;86;74m▄[m[38;2;214;252;186;48;2;0;0;0m█[m[38;2;214;252;186;48;2;0;0;0m█[m[38;2;215;252;186;48;2;0;0;0m█[m[38;2;215;252;186;48;2;0;0;0m█[m[38;2;216;252;185;48;2;0;0;0m█[m[38;2;216;252;185;48;2;0;0;0m█[m[38;2;216;252;185;48;2;0;0;0m█[m[38;2;0;0;0;48;2;49;49;49m [m[38;2;217;252;185;48;2;0;0;0m█[m[38;2;217;252;185;48;2;0;0;0m█[m[38;2;217;252;185;48;2;0;0;0m█[m[38;2;218;251;184;48;2;0;0;0m█[m[38;2;218;252;184;48;2;0;0;0m█[m[38;2;219;252;184;48;2;0;0;0m█[m[38;2;218;252;184;48;2;0;0;0m█[m[38;2;219;251;184;48;2;0;0;0m█[m[38;2;219;252;183;48;2;0;0;0m█[m[38;2;219;251;183;48;2;0;0;0m█[m[38;2;219;249;182;48;2;0;0;0m█[m[38;2;222;223;179;48;2;0;0;0m█[m[38;2;235;167;184;48;2;0;0;0m█[m[38;2;240;170;192;48;2;0;0;0m█[m[38;2;242;174;198;48;2;0;0;0m█[m[38;2;244;176;201;48;2;0;0;0m█[m[38;2;245;178;205;48;2;0;0;0m█[m[38;2;247;179;207;48;2;0;0;0m█[m[38;2;249;180;211;48;2;0;0;0m█[m[38;2;251;182;214;48;2;0;0;0m█[m[38;2;252;183;216;48;2;0;0;0m█[m[38;2;254;183;218;48;2;0;0;0m█[m[38;2;254;184;219;48;2;0;0;0m█[m[38;2;253;183;218;48;2;61;44;52m▀[m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;99;121;90;48;2;0;0;0m▀[m[38;2;210;240;187;48;2;0;0;0m█[m[38;2;231;208;188;48;2;0;0;0m█[m[38;2;201;246;184;48;2;0;0;0m█[m[38;2;205;253;189;48;2;133;157;123m▄[m[38;2;207;252;188;48;2;0;0;0m█[m[38;2;207;252;188;48;2;0;0;0m█[m[38;2;208;252;188;48;2;0;0;0m█[m[38;2;208;252;188;48;2;0;0;0m█[m[38;2;208;252;188;48;2;0;0;0m█[m[38;2;209;252;188;48;2;0;0;0m█[m[38;2;209;252;188;48;2;0;0;0m█[m[38;2;198;238;177;48;2;0;0;0m█[m[38;2;211;250;187;48;2;0;0;0m█[m[38;2;237;200;187;48;2;0;0;0m█[m[38;2;218;238;187;48;2;0;0;0m█[m[38;2;211;252;187;48;2;0;0;0m█[m[38;2;211;252;187;48;2;0;0;0m█[m[38;2;211;252;187;48;2;0;0;0m█[m[38;2;212;252;186;48;2;0;0;0m█[m[38;2;212;252;186;48;2;0;0;0m█[m[38;2;212;251;185;48;2;0;0;0m█[m[38;2;211;249;184;48;2;0;0;0m█[m[38;2;214;190;171;48;2;0;0;0m█[m[38;2;227;160;177;48;2;0;0;0m█[m[38;2;235;168;189;48;2;0;0;0m█[m[38;2;242;173;198;48;2;0;0;0m█[m[38;2;247;175;204;48;2;0;0;0m█[m[38;2;250;177;209;48;2;0;0;0m█[m[38;2;251;178;211;48;2;0;0;0m█[m[38;2;252;179;213;48;2;0;0;0m█[m[38;2;252;178;214;48;2;0;0;0m█[m[38;2;253;178;215;48;2;0;0;0m█[m[38;2;254;179;215;48;2;158;109;133m▀[m[38;2;253;179;215;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;16;11;13m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;100;126;95;48;2;52;67;51m▀[m[38;2;199;252;191;48;2;0;0;0m█[m[38;2;198;253;191;48;2;0;0;0m█[m[38;2;199;253;191;48;2;0;0;0m█[m[38;2;199;253;191;48;2;0;0;0m█[m[38;2;199;253;191;48;2;0;0;0m█[m[38;2;214;212;165;48;2;0;0;0m█[m[38;2;227;177;142;48;2;0;0;0m█[m[38;2;203;245;185;48;2;0;0;0m█[m[38;2;201;253;190;48;2;0;0;0m█[m[38;2;202;253;190;48;2;0;0;0m█[m[38;2;202;253;190;48;2;0;0;0m█[m[38;2;202;253;190;48;2;0;0;0m█[m[38;2;203;253;190;48;2;0;0;0m█[m[38;2;203;253;189;48;2;0;0;0m█[m[38;2;204;253;190;48;2;0;0;0m█[m[38;2;204;253;189;48;2;0;0;0m█[m[38;2;204;252;189;48;2;0;0;0m█[m[38;2;205;253;189;48;2;0;0;0m█[m[38;2;204;252;188;48;2;0;0;0m█[m[38;2;204;249;187;48;2;0;0;0m█[m[38;2;204;201;172;48;2;0;0;0m█[m[38;2;216;157;170;48;2;0;0;0m█[m[38;2;229;166;185;48;2;0;0;0m█[m[38;2;238;171;195;48;2;0;0;0m█[m[38;2;244;174;203;48;2;0;0;0m█[m[38;2;249;176;208;48;2;0;0;0m█[m[38;2;252;176;211;48;2;0;0;0m█[m[38;2;253;176;213;48;2;0;0;0m█[m[38;2;254;176;215;48;2;0;0;0m█[m[38;2;254;175;214;48;2;0;0;0m█[m[38;2;125;86;105;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;190;253;193;48;2;0;0;0m█[m[38;2;190;254;194;48;2;0;0;0m█[m[38;2;191;254;193;48;2;0;0;0m█[m[38;2;192;253;193;48;2;0;0;0m█[m[38;2;192;253;193;48;2;0;0;0m█[m[38;2;192;253;192;48;2;0;0;0m█[m[38;2;193;253;193;48;2;0;0;0m█[m[38;2;193;253;192;48;2;0;0;0m█[m[38;2;194;253;192;48;2;0;0;0m█[m[38;2;195;253;192;48;2;0;0;0m█[m[38;2;195;253;192;48;2;0;0;0m█[m[38;2;195;253;192;48;2;0;0;0m█[m[38;2;195;253;192;48;2;0;0;0m█[m[38;2;196;253;192;48;2;0;0;0m█[m[38;2;196;253;192;48;2;0;0;0m█[m[38;2;196;253;191;48;2;0;0;0m█[m[38;2;197;253;191;48;2;0;0;0m█[m[38;2;197;253;191;48;2;0;0;0m█[m[38;2;196;251;190;48;2;0;0;0m█[m[38;2;195;247;188;48;2;0;0;0m█[m[38;2;208;156;169;48;2;0;0;0m█[m[38;2;220;162;181;48;2;0;0;0m█[m[38;2;232;168;192;48;2;0;0;0m█[m[38;2;240;172;200;48;2;0;0;0m█[m[38;2;246;174;206;48;2;0;0;0m█[m[38;2;250;175;210;48;2;0;0;0m█[m[38;2;253;175;213;48;2;0;0;0m█[m[38;2;254;175;214;48;2;0;0;0m█[m[38;2;254;173;213;48;2;0;0;0m█[m[38;2;182;123;152;48;2;127;84;105m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;182;254;196;48;2;0;0;0m█[m[38;2;183;254;195;48;2;0;0;0m█[m[38;2;183;254;196;48;2;0;0;0m█[m[38;2;184;254;195;48;2;0;0;0m█[m[38;2;185;254;195;48;2;0;0;0m█[m[38;2;185;254;195;48;2;0;0;0m█[m[38;2;185;253;195;48;2;0;0;0m█[m[38;2;186;254;195;48;2;0;0;0m█[m[38;2;186;254;195;48;2;0;0;0m█[m[38;2;187;254;194;48;2;0;0;0m█[m[38;2;187;254;194;48;2;0;0;0m█[m[38;2;188;254;194;48;2;0;0;0m█[m[38;2;188;253;194;48;2;0;0;0m█[m[38;2;189;254;194;48;2;0;0;0m█[m[38;2;189;253;194;48;2;0;0;0m█[m[38;2;189;254;194;48;2;0;0;0m█[m[38;2;189;253;194;48;2;0;0;0m█[m[38;2;189;253;193;48;2;0;0;0m█[m[38;2;189;250;192;48;2;0;0;0m█[m[38;2;189;242;189;48;2;0;0;0m█[m[38;2;221;163;186;48;2;0;0;0m█[m[38;2;228;166;191;48;2;0;0;0m█[m[38;2;235;169;198;48;2;0;0;0m█[m[38;2;243;171;204;48;2;0;0;0m█[m[38;2;248;173;208;48;2;0;0;0m█[m[38;2;251;173;211;48;2;0;0;0m█[m[38;2;253;173;213;48;2;0;0;0m█[m[38;2;254;172;213;48;2;0;0;0m█[m[38;2;255;169;211;48;2;71;47;59m▀[m[38;2;0;0;0;48;2;2;1;1m [m[38;2;176;101;134;48;2;21;11;16m▄[m[38;2;251;150;195;48;2;0;0;0m█[m[38;2;253;160;203;48;2;0;0;0m█[m[38;2;235;153;192;48;2;0;0;0m█[m[38;2;159;106;133;48;2;0;0;0m▄[m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;175;254;198;48;2;0;0;0m█[m[38;2;176;254;198;48;2;0;0;0m█[m[38;2;176;254;197;48;2;0;0;0m█[m[38;2;177;254;197;48;2;0;0;0m█[m[38;2;177;254;197;48;2;0;0;0m█[m[38;2;178;254;197;48;2;0;0;0m█[m[38;2;178;254;197;48;2;0;0;0m█[m[38;2;179;254;197;48;2;0;0;0m█[m[38;2;179;254;197;48;2;0;0;0m█[m[38;2;180;254;196;48;2;0;0;0m█[m[38;2;180;254;196;48;2;0;0;0m█[m[38;2;181;254;196;48;2;0;0;0m█[m[38;2;181;254;196;48;2;0;0;0m█[m[38;2;182;254;196;48;2;0;0;0m█[m[38;2;182;254;196;48;2;0;0;0m█[m[38;2;182;254;195;48;2;0;0;0m█[m[38;2;182;254;195;48;2;0;0;0m█[m[38;2;182;253;195;48;2;0;0;0m█[m[38;2;181;251;193;48;2;0;0;0m█[m[38;2;177;242;187;48;2;0;0;0m█[m[38;2;136;96;114;48;2;0;0;0m▀[m[38;2;239;170;202;48;2;0;0;0m▀[m[38;2;243;170;205;48;2;102;70;86m▀[m[38;2;226;155;189;48;2;0;0;0m█[m[38;2;251;172;210;48;2;167;113;140m▀[m[38;2;253;172;212;48;2;11;7;9m▀[m[38;2;252;169;210;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;9;6;8m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;180;112;143;48;2;96;63;78m▀[m[38;2;250;165;204;48;2;0;0;0m█[m[38;2;251;169;208;48;2;0;0;0m█[m[38;2;253;171;212;48;2;0;0;0m█[m[38;2;248;167;207;48;2;181;122;151m▀[m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;3;5;4m [m[38;2;168;255;199;48;2;0;0;0m█[m[38;2;169;254;199;48;2;0;0;0m█[m[38;2;169;254;199;48;2;0;0;0m█[m[38;2;170;254;199;48;2;0;0;0m█[m[38;2;170;254;198;48;2;0;0;0m█[m[38;2;170;254;198;48;2;0;0;0m█[m[38;2;171;254;198;48;2;0;0;0m█[m[38;2;171;254;199;48;2;0;0;0m█[m[38;2;172;254;198;48;2;0;0;0m█[m[38;2;172;254;198;48;2;0;0;0m█[m[38;2;173;254;198;48;2;0;0;0m█[m[38;2;174;254;198;48;2;0;0;0m█[m[38;2;174;254;198;48;2;0;0;0m█[m[38;2;175;254;197;48;2;0;0;0m█[m[38;2;175;254;197;48;2;0;0;0m█[m[38;2;176;254;198;48;2;0;0;0m█[m[38;2;176;254;197;48;2;0;0;0m█[m[38;2;175;254;197;48;2;0;0;0m█[m[38;2;175;252;195;48;2;0;0;0m█[m[38;2;173;248;193;48;2;0;0;0m█[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;87;59;73;48;2;0;0;0m▀[m[38;2;253;171;211;48;2;0;0;0m▀[m[38;2;215;144;180;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;3;4;3m [m[38;2;163;255;201;48;2;0;0;0m█[m[38;2;164;254;200;48;2;0;0;0m█[m[38;2;164;254;200;48;2;0;0;0m█[m[38;2;163;254;199;48;2;0;0;0m█[m[38;2;148;230;180;48;2;0;0;0m█[m[38;2;165;254;199;48;2;2;3;2m▀[m[38;2;160;248;194;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;2;4;3m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;1;2;1m [m[38;2;159;240;188;48;2;0;0;0m▀[m[38;2;169;254;198;48;2;0;0;0m▀[m[38;2;152;230;180;48;2;0;0;0m█[m[38;2;169;254;199;48;2;0;0;0m█[m[38;2;170;254;199;48;2;0;0;0m█[m[38;2;170;254;199;48;2;0;0;0m█[m[38;2;170;254;199;48;2;0;0;0m█[m[38;2;170;253;198;48;2;0;0;0m█[m[38;2;168;251;197;48;2;0;0;0m█[m[38;2;0;0;0;48;2;0;1;1m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;77;122;96;48;2;0;0;0m▀[m[38;2;151;240;190;48;2;0;0;0m▀[m[38;2;72;114;89;48;2;0;0;0m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;5;8;6m [m[38;2;162;249;195;48;2;0;0;0m▀[m[38;2;166;255;200;48;2;25;39;31m▀[m[38;2;163;251;197;48;2;0;0;0m█[m[38;2;165;254;200;48;2;0;0;0m█[m[38;2;165;252;199;48;2;66;103;81m▀[m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m[38;2;0;0;0;48;2;0;0;0m [m
`

	if !reflect.DeepEqual(result, expected) {
		a := os.WriteFile("/tmp/dat1", []byte(result), 0644)
		println(a)
		t.Errorf("Result and expected does not match = %v, want %v", result, expected)
	}

	// used for quick testing
	// fmt.Println(result)
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return png.Decode(f)
}
