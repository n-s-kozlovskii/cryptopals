import scala.runtime.Arrays
import java.util.Base64
import java.nio.charset.StandardCharsets

class ReaptedXor {
  def countDifferentBits(a: Byte, b: Byte): Int = {
    var count: Int = 0

    for i <- 0 to 8 do {
      val a1 = a >> i
      val a2 = a1 & 1
      val b1 = b >> i
      val b2 = b1 & 1

      if (b2 != a2) {
        count += 1
      }
      //    println(f"i=$i;count=$count\n\ta1:${a1.toInt.toBinaryString}\n\tb1:${b1.toInt.toBinaryString}")
    }

    count
  }

  def hamming(a: Array[Byte], b: Array[Byte]): Int = {
    assert(a.length == b.length)
    var distance = 0
    for (e, g) <- a zip b do {
      distance += countDifferentBits(e, g)
    }

    distance
  }

  def estimateKeySize(a: Array[Byte]): Double = {
    assert(a.length % 2 == 0)
    val keySize = a.length / 2
    val (firstHalf, secondHalf) = a.splitAt(keySize)

    hamming(firstHalf, secondHalf).toDouble / keySize
  }

  def findSmallestKeySizes(a: Array[Byte]): Map[Int, Double] = {
    assert(a.length > 100)

    (2 to 50).map { i =>
      i -> estimateKeySize(a.slice(0, i * 2))
    }.toMap
  }

  def fromFile(): Seq[Int] = {
    val source = scala.io.Source.fromFile("../../testdata/task6_input.txt")
    val lines =
      try source.mkString
      finally source.close()
    val data = Base64.getDecoder().decode(lines)
    val map = findSmallestKeySizes(data)
    println(map)
    map.toSeq.sortWith(_._2 < _._2).map(_._1)
  }
}
