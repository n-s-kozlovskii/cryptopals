// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class MySuite extends munit.FunSuite {
  val testObject = new ReaptedXor()

  test("can count number of differing bit") {

    // 10 -> 1010
    // 14 -> 1110
    // 13 -> 1111
    assertEquals(testObject.countDifferentBits(10, 14), 1)
    assertEquals(testObject.countDifferentBits(8, 15), 3)
  }

  test("can find hamming distance") {
    val a = "this is a test".getBytes()
    val b = "wokka wokka!!!".getBytes()

    assertEquals(testObject.hamming(a,b), 37)
  }

  test("can estimate keysize") {
    val a = "this is a testwokka wokka!!!".getBytes()

    assertEquals(testObject.estimateKeySize(a), 37.0/14)
  }

  test("can find smallest keysize") {
    val res = testObject.fromFile()
    assertEquals(res, List(1,2,3))
  }
}
