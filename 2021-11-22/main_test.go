package november222021

import (
	"testing"
)

var benchListOfStuff = []string{"eat", "tea", "ten", "poop", "net", "ate"}
var benchListOfStuffLong = []string{"lbzgb", "icm", "ajww", "thct", "uaxh", "kqfda", "plsjf", "cxo", "ffrswxp", "dnjob", "snvlgt", "mape", "qleqyh", "zrywjjp", "zpfrfe", "motafe", "hsbzrj", "awnwe", "rbemfd", "dcekx", "akjqzlc", "tmt", "coanaty", "inkar", "kjyi", "jrs", "ctnsw", "nsgru", "svma", "zfzbs", "ojifqg", "snwtk", "mvoigl", "pbuop", "dkupd", "mervja", "zln", "xyeuc", "ksx", "gyra", "mbtvksj", "jzalb", "zsy", "geudt", "zqmdqi", "cohgho", "gseycj", "jhynu", "njjhhju", "rusqfgq", "mkp", "vkur", "pifv", "zrgbmya", "kctzk", "kzivabj", "kxvbw", "vbq", "gexyalb", "djsgpn", "cwfkdi", "ibuuff", "owditsk", "oqjmqr", "ictoji", "xyes", "zyfroro", "mbndrzn", "nrwcjpm", "dtjmha", "orsufum", "psv", "zhb", "myyt", "jvgwffb", "ggcnqba", "reun", "zjqxmzo", "arlut", "ygm", "vyb", "ddvo", "ifsfgpy", "kmxiube", "tndtj", "yrrded", "iylpruc", "iogjhye", "wbtcmlf", "dgxqwp", "wvgq", "zclvc", "asj", "dsy", "ofkkey", "qkkhqg", "pnbpbg", "mluid", "ummp", "hcs", "mjjxz", "aii", "nbakqs", "qpo", "gnczgac", "ainlql", "baatly", "daopo", "fok", "iexsfzx", "rlcztx", "djjfuy", "hrcovgp", "vlg", "xalgqar", "nebzbf", "lhxkz", "fnav", "ayyqwzk", "qfbuc", "njywrn", "gkkldt", "nyocsfk", "ohsvvx", "azwexe", "haqu", "daaazlr", "onx", "payoss", "cnctugz", "mctoz", "pyna", "phid", "akua", "mbdt", "tcof", "spqkxs", "efzap", "jzldaue", "hitg", "vbrqpq", "arpx", "tpv", "npdgerw", "hgcmd", "lit", "qwluecg", "czx", "brmgxq", "exouabu", "qripjy", "yqfs", "fubv", "dhtaknj", "qxcqk", "ift", "xejtun", "bfqqusx", "odpo", "vauk", "wwwt", "dujhiqe", "bxzvqzl", "wyqos", "blnaip", "yxd", "hcbc", "rvd", "sfowl", "rcnpa", "wvkkod", "jvcugm", "tnmuq", "cvfoh", "tut", "wjrgi", "jotcvw", "jkvtp", "vvtynxf", "jpcd", "mixvd", "zfh", "nmkguot", "uzjuuzk", "wqsr", "ovwaye", "kzpe", "rsokaak", "kawl", "vcjy", "cwvx", "rnsfu", "cfs", "yszwtbp", "vkb", "rkvuduy", "sgd", "nqxe", "jnkxlk", "kwwu", "qsfz", "xvjtsnz", "mkcsist", "duz", "oleziji", "mapkvat", "jzyev", "msidsj", "kovp", "qsnynsl", "vmwzlbk", "ziypt", "zlf", "huhsw", "dgpy", "srckct", "dcuo", "npmv", "zdiy", "ouqzmjg", "oxytsbf", "aqtx", "kvit", "bhgi", "arc", "bbl", "mhef", "wcnjq", "kslxkk", "vyx", "ocd", "tiu", "jevakgz", "altlhtu", "hjox", "mrnwkdo", "nxzmdik", "jpukql", "xahzfs", "szrqgn", "dbebys", "iwckgfx", "val", "kap", "qquqnzx", "fbig", "pkmldqo", "jks", "jfllbl", "uehz", "xquvapl", "pkp", "xdpizdh", "gjsj", "bqa", "fyq", "yzd", "shiiwpu", "aqbhq", "zocikj", "qbh", "iwq", "zdsyrmm", "eyrr", "ueqcu", "dmne", "iktxrin", "szjai", "xbrd", "lmciv", "wluwp", "dpi", "zuui", "fxrgxb", "mbje", "vzwc", "xyxloe", "eaketw", "udx", "dutrnr", "hmslaos", "uvszs", "hkange", "zczewg", "rduc", "wptqaq", "tyv", "pvpkh", "gwhxbh", "sac", "aav", "srdtx", "dgnmc", "mofzvwi", "msn", "kzliun", "scrakq", "rsfk", "aist", "mupn", "mva", "vez", "ymxcwr", "luxbdm", "yhecgct", "ubglun", "yyo", "amr", "xygzwgl", "dnnmawe", "syzbjg", "aeqx", "vlh", "eeizrbu", "spxmq", "mdopo", "nqvns", "kke", "fulxs", "awwjt", "zbrw", "kvdea", "nupka", "borkco", "yawyay", "edjztrp", "tlj", "jynub", "yuretw", "toopui", "fonfw", "xmpg", "uami", "abzm", "ulh", "kkv", "vnxeum", "unfi", "pjh", "cdmc", "ryregnc", "cep", "nim", "ocnrt", "hzshv", "iif", "qrwxv", "mhxnhuq", "wvaqhnq", "rqs", "fjldkm", "jvvofl", "vss", "ayk", "kri", "xeovecl", "uundnjg", "nlvgp", "tbpzgjz", "dkac", "eotkif", "ofrksn", "kexqbqz", "wdjnu", "selb", "yoz", "jyzuy", "iivqnia", "tai", "fjxvlk", "qbjaj", "gvwvwgc", "iuyxyh", "phnv", "qqh", "nzdrit", "uovpog", "tittv", "aedtmb", "keek", "gqi"}

func bench(list []string, f func([]string) [][]string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(list)
	}
}
func BenchmarkGroupByStringsLong(b *testing.B) {
	bench(benchListOfStuffLong, groupByAnagramsButtStrings, b)
}
func BenchmarkGroupByNumbersLong(b *testing.B) { bench(benchListOfStuffLong, groupByAnagrams, b) }
func BenchmarkGroupByStringsShort(b *testing.B) {
	bench(benchListOfStuff, groupByAnagramsButtStrings, b)
}
func BenchmarkGroupByNumbersShort(b *testing.B) { bench(benchListOfStuff, groupByAnagrams, b) }
