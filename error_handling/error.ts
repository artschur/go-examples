async function foo() {
  throw new Error('An error occurred in foo');
}

async function bar() {
  throw new Error('An error occurred in bar');
}

async function main() {
  var fooVal, barval;
  try {
    const fooVal = await foo();
  } catch (error) {
    return;
  }
  try {
    const barval = await bar();
  } catch (error) {
    return;
  }

  return fooVal + barval;
}
