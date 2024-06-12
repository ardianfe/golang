package maps

import (
	"testing"
)

func TestMap(t *testing.T) {
	buah := map[string]int{
		"anggur": 2000,
		"apel":   5000,
		"melon":  3000,
	}

	if len(buah) != 3 {
		t.Errorf("Harus 3 item tapi terdapat %d", len(buah))
	}

	if buah["melon"] != 3000{
		t.Errorf("Harga harusnya 3000 bukan %d",buah["melon"])
	}

}

func TestManipulatinMaps(t *testing.T)  {
	makan := ManipulatingMaps()

	if len(makan)!=2{
		t.Errorf("Harus 2 item, tapi terdapat %d",len(makan))
	}

	if _, cek:=makan["bakso"];cek{
		t.Errorf("Bakso harusnya dihapus")
	}
}
