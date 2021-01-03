package binson

import (
    "testing"
    "reflect"
    "encoding/hex"
    "github.com/stretchr/testify/assert"
)

func ByteEqual(t *testing.T, got []byte, wantString string){
    want, err := hex.DecodeString(wantString)
    //if err != nil || bytes.Compare(got, want) != 0 {
    if err != nil || !reflect.DeepEqual(got, want){
        t.Errorf("Failed\nGot:  %s\nWant: %s", hex.EncodeToString(got), wantString)
    }
}

func TestBinson(t *testing.T) {
    want, _ := hex.DecodeString("4041")
    b := NewBinson()
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")

    want, _ = hex.DecodeString("40140161100441")
    b = NewBinson().
        PutInt("a", 4)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")

    want, _ = hex.DecodeString("4014016214044772697341")
    b = NewBinson().
        PutString("b", "Gris")
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    want, _ = hex.DecodeString("4014016340140164140348656a4141")
    b = NewBinson().
        PutBinson("c", NewBinson().
            PutString("d", "Hej"))
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    want, _ = hex.DecodeString("4014016142424310024341")
    b = NewBinson().
        PutArray("a", NewBinsonArray().
            PutArray(NewBinsonArray()).
            PutInt(2) )
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    want, _ = hex.DecodeString("40140161180301020341")
    b = NewBinson().
        PutBytes("a", []byte{1, 2, 3})
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    want, _ = hex.DecodeString("40140161441401624541")
    b = NewBinson().
        PutBool("a", true).
        PutBool("b", false)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
	
	want, _ = hex.DecodeString("4014016146e17a14ae4701374041")
    b = NewBinson().
        PutFloat("a", 23.005)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")

    want, _ = hex.DecodeString("401401611004140162140467696769140163404114016442431401651803010203140166441401674614ae47e17a543e4041")
    b = NewBinson().
        PutInt("a", 4).
        PutString("b", "gigi").
        PutBinson("c", NewBinson()).
        PutArray("d", NewBinsonArray()).
        PutBytes("e", []byte{1, 2, 3}).
        PutBool("f", true).
		PutFloat("g", 30.33)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    b = NewBinson().
        Put("a", 4).
        Put("b", "gigi").
        Put("c", NewBinson()).
        Put("d", NewBinsonArray()).
        Put("e", []byte{1, 2, 3}).
        Put("f", true).
		Put("g", 30.33)
    assert.Equal(t, want, b.ToBytes(), "Bytes do not match")
    
    want, _ = hex.DecodeString("421004140467696769404142431803010203454614ae47e17a543e4043")
    a := NewBinsonArray().
        PutInt(4).
        PutString("gigi").
        PutBinson(NewBinson()).
        PutArray(NewBinsonArray()).
        PutBytes([]byte{1, 2, 3}).
        PutBool(false).
		PutFloat(30.33)
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")

    a = NewBinsonArray().
        Put(4).
        Put("gigi").
        Put(NewBinson()).
        Put(NewBinsonArray()).
        Put([]byte{1, 2, 3}).
        Put(false).
		Put(30.33)
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")

    b = NewBinson().
        Put("a","g").
        Put("b","h")
    assert.Equal(t, b.FieldNames(), []string{"a", "b"}, "Keys do not match")    
    assert.True(t, b.ContainsKey("a"), "Key do not exist")
    assert.True(t, b.ContainsKey("b"), "Key do not exist")
    assert.False(t, b.ContainsKey("c"), "Key should not exist")
    b.Remove("a")
    assert.Equal(t, b.FieldNames(), []string{"b"}, "Keys do not match")
    assert.False(t, b.ContainsKey("a"), "Key should not exist")
    assert.True(t, b.ContainsKey("b"), "Key do not exist")
    
    a = NewBinsonArray().
        Put("a").
        Put("b").
        Put("c")
    assert.Equal(t, 3, a.Size(), "Wrong length")
    want, _ = hex.DecodeString("4214016114016214016343")
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")
    a.Remove(1)
    assert.Equal(t, 2, a.Size(), "Wrong length")
    want, _ = hex.DecodeString("4214016114016343")
    assert.Equal(t, want, a.ToBytes(), "Bytes do not match")
}


/*
public class Examples {
	public static void main(String[] args) {
		Binson ex1 = new Binson().put("cid", 4);
		System.out.println("Example 1:");
		System.out.println("  {cid=4;}");
		System.out.println("  " + hex(ex1.toBytes()));
		
		Binson ex2 = new Binson();
		System.out.println("Example 2, empty object:");
		System.out.println("  {}");
		System.out.println("  " + hex(ex2.toBytes()));
		
		Binson ex3 = new Binson().put("a", new Binson().put("b", 2));
		System.out.println("Example 3, nested object:");
		System.out.println("  {a={b=2;};");
		System.out.println("  " + hex(ex3.toBytes()));
		
		Binson ex4 = new Binson()
				.put("a", 1)
				.put("b", new Binson().put("c", 3))
				.put("d", 4);
		System.out.println("Example 4, object field between integer fields:");
		System.out.println("  {a=1; b={c=3;}; d=4}");
		System.out.println("  " + hex(ex4.toBytes()));
		
		Binson ex5 = new Binson()
				.put("a", new BinsonArray().add(1).add("hello"));
		System.out.println("Example 5, array");
		System.out.println("  {a=[1, \"hello\"];}");
		System.out.println("  " + hex(ex5.toBytes()));
		
		Binson ex6 = new Binson()
				.put("a", 1)
				.put("b", new BinsonArray().add(10).add(20))
				.put("c", 3);
		System.out.println("ex6, array");
		System.out.println("  {a=1; b=[10,20]; c=3}");
		System.out.println("  " + hex(ex6.toBytes()));

		Binson ex7 = new Binson()
				.put("a", 1)
				.put("b", new BinsonArray().add(10).add(new BinsonArray().add(100).add(101)).add(20))
				.put("c", 3);
		System.out.println("ex7, array");
		System.out.println("  {a=1; b=[10, [100, 101], 20]; c=3}");
		System.out.println("  " + hex(ex7.toBytes()));
		
		Binson ex8 = new Binson()
				.put("a", 1)
				.put("b", -1)
				.put("c", 250)
				.put("d", Integer.MAX_VALUE)
				.put("f", Long.MAX_VALUE);
		System.out.println("ex8, array");
		System.out.println("  {a=1; b=-1; c=250; d=Integer.MAX_VALUE, f=Long.MAX_VALUE");
		System.out.println("  " + hex(ex8.toBytes()));
		
		Binson ex9 = new Binson()
				.put("aaaa", 250);
		System.out.println("ex9, int value = 250");
		System.out.println("  {aaaa=250}");
		System.out.println("  " + hex(ex9.toBytes()));
		
		Binson ex10 = new Binson()
                .put("aaaa", "bbb");
        System.out.println("ex10, short string value");
        System.out.println("  {aaaa=\"bbb\"}");
        System.out.println("  " + hex(ex10.toBytes()));
        
        Binson ex11 = new Binson()
                .put("aa", new byte[]{5, 5, 5});
        System.out.println("ex11, short bytes value");
        System.out.println("  {aa=0x050505;}");
        System.out.println("  " + hex(ex11.toBytes()));
	}
	
	private static String hex(byte[] bytes) {
		return JsonOutput.bytesToHex("", bytes);
	}
}
*/

/*
Output (2016-01-06):
Example 1:
  {cid=4;}
  401403636964100441
Example 2, empty object:
  {}
  4041
Example 3, nested object:
  {a={b=2;};
  401401614014016210024141
Example 4, object field between integer fields:
  {a=1; b={c=3;}; d=4}
  40140161100114016240140163100341140164100441
Example 5, array
  {a=[1, "hello"];}
  40140161421001140568656c6c6f4341
ex6, array
  {a=1; b=[10,20]; c=3}
  40140161100114016242100a101443140163100341
ex7, array
  {a=1; b=[10, [100, 101], 20]; c=3}
  40140161100114016242100a421064106543101443140163100341
ex8, array
  {a=1; b=-1; c=250; d=Integer.MAX_VALUE, f=Long.MAX_VALUE
  40140161100114016210ff14016311fa0014016412ffffff7f14016613ffffffffffffff7f41
ex9, int value = 250
  {aaaa=250}
  4014046161616111fa0041
ex10, short string value
  {aaaa="bbb"}
  40140461616161140362626241
ex11, short bytes value
  {aa=0x050505;}
  4014026161180305050541
*/