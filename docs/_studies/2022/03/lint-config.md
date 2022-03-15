---
title: ESLint Config Improvement Proposal
date: 2022-03-15
excerpt: 'A review of the most recent updates to the GPA Lab ESLint config.'
---

## General Config Changes

1 - Deprecated the Prettier configuration since we no longer use Prettier to format our code. This allows us to also removed the related packages (ie. `prettier`, `eslint-plugin-prettier`, and `@gpalab/prettier-config`) thereby shaving down our dependency count a bit.

<hr/>

2 - Add an override for Webpack config files by targeting all files named using the following pattern `webpack.*.js`. This override disables the [unpublished require statements](https://github.com/mysticatea/eslint-plugin-node/blob/master/docs/rules/no-unpublished-require.md) rule for Webpack configs. This is helpful because often webpack plugins are required by these configs but installed as dev dependencies, which results in linting errors that must be manually addressed in the project ESLint config file. This is a common enough occurrence that the override seems worthwhile.

## New to ESLint

Several new rules and options were added to ESLint since the last time the configuration was updated (ie. ESLint v7.32.0 => v8.11.0).

1 - [no-unused-private-class-members](https://eslint.org/docs/rules/no-unused-private-class-members) new rule introduced in v8.1.0. Using the rationale that unused code adds unnecessary noise and should be removed, this rule indicates the presence of private methods that are declared but never utilized. Enabled at the level `warn` for the time being, it is unlikely to have a noticeable impact since we don't use private methods at the moment.

```jsx
class MyClass extends Component {
  // ❌ This method would be flagged because it is never used.
  #badMethod() {
    console.log('Bad');
  }

  // ✅ This method is okay because it is used in the 'finalMethod' below.
  #goodMethod() {
    console.log('Good');
  }

  finalMethod() {
    this.#goodMethod();
  }
}
```

<hr/>

2 - Set the new [destructuredArrayIgnorePattern](https://eslint.org/docs/rules/no-unused-vars#destructuredarrayignorepattern) property on `no-unused-vars` to `^_`. This allows us to indicate unused array elements when destructuring, by adding a leading underscore, and thereby avoiding unused variable errors.

```js
// ✅ This is okay because the unused item '_b' if prefixed by '_'.
const [a, _b, c] = ['a', 'b', 'c'];
console.log(a + c);

// ❌ This would throw an error since the unused variable is not preceded by an underscore.
const [d, unused, f] = ['d', 'e', 'f'];
console.log(d + f);
```

<hr/>

3 - [prefer-object-has-own](https://eslint.org/docs/rules/prefer-object-has-own) new rule introduced in v8.5.0. This rule favors the use of `Object.hasOwn()` over `Object.prototype.hasOwnProperty.call()`, neither of which is a pattern that we typically use. Furthermore, this rule is only relevant in an ES2020 context which may or may not be true for any given Lab project. For these reasons this rule is set to `off`.

<hr/>

4 - Set the new [enforceForClassFields](https://eslint.org/docs/rules/class-methods-use-this#enforceforclassfields) property on `class-methods-use-this` and the [onlyOneSimpleParam](https://eslint.org/docs/rules/no-confusing-arrow) property on `no-confusing-arrow` to `false`. These options seem to have little relevance for our code so these choices simply explicitly select the current default behavior.

## New to ESLint Plugin React

Several new rules and options were added to ESLint React plugin since the last time the configuration was updated (ie. ESLint Plugin React 7.26.0 => 7.29.3).

1 - [react/no-arrow-function-lifecycle](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/no-arrow-function-lifecycle.md) new rule introduced in v7.27.0 of `eslint-plugin-react`. Prevents the declaration of React class lifecycle methods as arrow functions, since doing so can be "conceptually less performant" and can break hot reloading.

```jsx
class MyComponent1 extends Component {
  // ✅ This method would not be flagged.
  render() {
    return <div />;
  }
}

class MyComponent2 extends Component {
  // ❌ This method would be flagged.
  render = () => {
    return <div />;
  };
}
```

<hr/>

2 - [react/no-invalid-html-attribute](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/no-invalid-html-attribute.md) new rule introduced in v7.27.0 of `eslint-plugin-react`. Certain HTML element (ex. `a`, `area`, `link`, `form`) have native `rel` attributes that accepts a [constrained list of values](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/rel). This rule prevents the developer from using an invalid `rel` value on these elements.

```html
<!-- ✅ Alternate is a valid rel value for 'a' elements. -->
<a href="https://example.com" rel="alternate">Test</a>

<!-- ❌ Canonical is not a valid rel value for 'a' elements. -->
<a href="https://example.com" rel="canonical">Test</a>
```

<hr/>

3 - [react/no-unused-class-component-methods](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/no-unused-class-component-methods.md) new rule introduced in v7.27.0 of `eslint-plugin-react`. It ensures that all methods declared in class components are used. Unused methods are often the result of incomplete refactors and if nothing else obscure useful code.

```jsx
class MyClass extends Component {
  // ❌ This method would be flagged because it is never used.
  badMethod() {
    console.log('Bad');
  }

  // ✅ This method is okay because it is used in the class's lifecycle below.
  goodMethod() {
    console.log('Good');
  }

  componentDidMount() {
    this.goodMethod();
  }

  render() {
    return null;
  }
}
```

<hr/>

4 - [react/hook-use-state](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/hook-use-state.md) new rule introduced in v7.29.0 of `eslint-plugin-react`. Enforce the `[value,setValue]` naming convention when destructuring the return from a `useState` call.

```jsx
const [color, setColor] = useState('#ffffff'); // ✅ Preferred syntax.
const [color, updateColor] = useState('#ffffff'); // ❌ Setter name does not follow convention.
const [color, setHex] = useState('#ffffff'); // ❌ Lack of symmetry between value and setter.
```

<hr/>

5 - [react/iframe-missing-sandbox](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/iframe-missing-sandbox.md) new rule introduced in v7.29.0 of `eslint-plugin-react`. Forces developers to follow security best practices and require a [sandbox property](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#attr-sandbox) on iframe elements. To apply all sandbox restriction, the property should be added with an empty value

```html
<!-- ✅ This iframe has the sandbox attribute correctly set. -->
<div>
  <iframe sandbox="allow-same-origin"></iframe>
</div>

<!-- ❌ This iframe is missing the sandbox attribute. -->
<div>
  <iframe></iframe>
</div>
```

<hr />

6 - Utilize the new `propElementValues` property on [react/jsx-curly-brace-presence](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/jsx-curly-brace-presence.md) to ensure that when passing a JSX element as a prop, it is always surrounded by curly brackets.

```jsx
// ✅ Element passed as prop should be wrapped in curly braces.

<MyComponent jsxProp={<div />} />
```

```jsx
// ❌ Valid JSX but confusing to read.

<MyComponent jsxProp=<div /> />
```

<hr/>

7 - Enable the new `warnOnDuplicates` property on [react/jsx-key](https://github.com/yannickcr/eslint-plugin-react/blob/master/docs/rules/jsx-key.md) to detect non-unique keys in an array.

```jsx
// ❌ With this option, such non-unique keys will be highlighted.
const elements = [
  <span key="123" />
  <span key="123" />
]
```

## Previously Omitted Rules:

Finally, several existing rules and options that had been unwittingly omitted in previous versions of this configuration were added.

1 - [jsx-a11y/autocomplete-valid](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y/blob/main/docs/rules/autocomplete-valid.md) - set to `error` to ensure that the [autocomplete attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes/autocomplete) is only used with appropriate input elements. (added in v6.3.0 of eslint-plugin-jsx-a11y)

```html
<!-- ✅ The autocomplete attribute on this input element matches the input type. -->
<input type="email" autocomplete="email" />

<!-- ❌ An email input should not accept a URL autocomplete value. -->
<input type="email" autocomplete="url" />
```

<hr/>

2 - Enable the [func-name-matching](https://eslint.org/docs/rules/func-name-matching) rule. In practice we have been already following this rule.

```js
// ✅ Variable assignment matches the function name.
const hello = function hello() {
  return 'hello';
};

// ❌ Variable assignment is different than the function name.
const hi = function hello() {
  return 'hello';
};
```

<hr/>

3 - Enable the [require-atomic-updates](https://eslint.org/docs/rules/require-atomic-updates) rule at the `warn` level. This rule handles cases where a variable is both read and reassigned by an asynchronous function. This can lead to subtle bugs where in the variable is reassigned and used prior to the completion of a required asynchronous action thereby returning an improper value.

```js
let result;

// ❌ With this function, the 'result' variable could be
// used in error before the awaited action has completed.
async function bar() {
  result = result + (await something);
}

// ✅ Here the async return is stored in a local variable so the 'result'
// variable cannot be accessed while the function is paused.
async function baz() {
  const tmp = doSomething(await somethingElse);
  result += tmp;
}
```

<hr/>

4 - Use a stricter configuration of the `no-shadow` rule by setting the [builtinGlobals](https://eslint.org/docs/rules/no-shadow#builtinglobals) property to `true`. This disables shadowing of global variables (ex. `Object` and `Array`).

```js
// ❌ With this option enabled, these declarations would throw an errors
// since we are attempting to overwrite a built-in global variable.
const Object = 'hello';

let Array = 1;
```

<hr/>

5 - Use a stricter configuration of the [no-inner-declarations](https://eslint.org/docs/rules/no-inner-declarations) rule disallowing variable declarations as well as function declarations. This rule aims to avoid issues where deeply nested `var` declarations are unintentionally hoisted to the program root. We typically use block bound `let` and `const` declarations so this may be a bit redundant, as the are not subject to hoisting.

<hr/>

6 - The missing ESLint import plugin rules `import/no-relative-packages` and `import/no-import-module-exports` were added to the config for the sake of completeness, but at the level `off` so they will have no impact.
