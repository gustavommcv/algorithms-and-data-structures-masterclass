import assert from "assert";
import factorial from "./factorial.js";

assert.strictEqual(factorial(0), 1, "result != 1");
assert.strictEqual(factorial(1), 1, "result != 1");
assert.strictEqual(factorial(7), 5040, "result != 5040");

console.log("All tests passed")
