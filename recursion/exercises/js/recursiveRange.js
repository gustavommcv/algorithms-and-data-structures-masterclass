export default function recursiveRange(number) {
  if (number < 1) {
    return 0;
  }

  return number + recursiveRange(number - 1);
}
