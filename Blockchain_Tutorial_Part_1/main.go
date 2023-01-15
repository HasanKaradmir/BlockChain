package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type BlockChain struct {
	// Bu struct blocklarin hepsini ifade etmek icin olusturuldu.
	blocks []*Block
	// Bu struct suan icin verimli calisak fakat block'lar olusmaya basladikca
	// daha da karmasik hale gelecektir.
}

type Block struct {
	Hash []byte
	// Hash'lenmis veriyi tutar.
	Data []byte
	// Asil veriyi tutar. Resim, dosya vb.
	PrevHash []byte
	// Bir onceki hash'e sahip son block hash'i, blocklari bir tur arka baglantili
	//liste gibi birbirine baglamamiza izin verir. Blockchain icerisinde olusturulan son bloga referans
	//verir ve biz aslinda 'Hash' ve 'Data'da olusanlarin uzerinde karmasik hale gelebilen block boyutu,
	//zaman damgasi ve bu hesaplamaya dahil olan birkac baska alan gibi seyler de eklenecek.

}

func (b *Block) DeriveHash() {
	// Bu fonksiyon bizim verdigimiz degerleri hashlemek icin kullaniliyor.
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Join, yeni bir bayt dilimi oluşturmak için öğeleri birleştirir.
	hash := sha256.Sum256(info)
	// Hash'leme islemini yapar.
	b.Hash = hash[:]
	// hash'lenmis veriyi degiskene atiyor.
}

func CreateBlock(data string, prevHash []byte) *Block {
	// Bu fonksiyon yeni bir block olusturur.
	block := &Block{[]byte{}, []byte(data), prevHash}
	// Block struct'ini block olarak gosterip bunlari degiskene atiyor.
	block.DeriveHash()
	//Block hash'leniyor.
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	// Bu fonksiyon blockchain'e yeni bir block olusturmak icin olusturuldu.
	prevBlock := chain.blocks[len(chain.blocks)-1]
	// blockchain'deki onceki block'u almak istiyoruz ve bunu zincirlenmis block'lari cagirarak
	// yababiriz ve ardindan blockchain block'larimizin uzunlugundan 1 cikararak bu block'u olustururuz.
	new := CreateBlock(data, prevBlock.Hash)
	// Burada data'da gecen CreateBlock fonksiyonu cagiriyoruz ve ardindan prevBlock hash'ini olarak gecerli blogu veriyoruz
	chain.blocks = append(chain.blocks, new)
	// Ardindan chain.block dizisine yeni (new) olusturulan blogu ekliyoruz.
}

// Blocklari olusturduk tamam ama ilk block ne olacak? Bunu da olusturuyoruz. Buna Genesis Block denir. Yani BASLANGIC BLOGU.

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
	// Tek yapmamiz gereken islem, block olusturmak.
}

func InitBlockChain() *BlockChain {
	// Bu fonksiyon blockchain baslatmak ve calistirmak icin olusturuluyor.
	return &BlockChain{[]*Block{Genesis()}}
	// Blockchain olusturmak icin elimizdeki ilk block'u (Genesis) veriyoruz ve boylece InitBlockChain'i calistiriyoruz.
}

// Tum bunlari yaptiktan sonra main'de blockchain'i calistirabiliriz.

func main() {
	chain := InitBlockChain()
	// Blockchain calistirma islevini bir degiskene atiyoruz.
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	// Burada da blockchain'i calistirdiktan sonra yeni block'lar ekliyoruz.
	for _, block := range chain.blocks {
		// Bu dongu blockchain'i gormek icin olusturuluyor.
		fmt.Printf("\nPrevious Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
