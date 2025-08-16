import assert from "assert";
import recursiveRange from "./recursiveRange.js";

assert.strictEqual(recursiveRange(6), 21, "result != 21");
assert.strictEqual(recursiveRange(10), 55, "result != 55");

console.log("All tests passed");
