import assert from "assert";
import productOfArray from "./productOfArray.js";

assert.strictEqual(productOfArray([1, 2, 3]), 6, "result != 6");
assert.strictEqual(productOfArray([1, 2, 3, 10]), 60, "result != 60");

console.log("All tests passed");
