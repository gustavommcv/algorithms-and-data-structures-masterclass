import assert from "assert";
import binarySearch from "./binarySearch.js";

assert.strictEqual(binarySearch([1, 2, 3, 4, 5], 3), 2, "result != 2");
assert.strictEqual(binarySearch([1, 2, 3, 4, 5], 7), -1, "result != -1");

console.log("All tests passed");
