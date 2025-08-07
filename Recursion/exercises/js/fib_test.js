import assert from "assert";
import fib from "./fib.js";

assert.strictEqual(fib(4), 3, "result != 3");
assert.strictEqual(fib(10), 55, "result != 55");
assert.strictEqual(fib(28), 317811, "result != 317811");
assert.strictEqual(fib(35), 9227465, "result != 9227465");

console.log("All tests passed");
