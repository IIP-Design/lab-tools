---
title: ESLint Config Improvement Proposal
date: 2022-03-15
excerpt: 'A review of the most recent updates to the GPA Lab ESLint config.'
---

Deprecated the Prettier configuration, removed the related packages (`prettier`, `eslint-plugin-prettier`, and `@gpalab/prettier-config`).

An override to allow unpublished require statements in Webpack config files (identified by the filenames `webpack.*.js`).

## New to ESLint

Several new rules and options were added to ESLint since the last time the configuration was updated (ie. ESLint v7.32.0 -> v8.11.0).

1 - [no-unused-private-class-members](https://eslint.org/docs/rules/no-unused-private-class-members) new rule introduced in v8.1.0. Indicates the presence of private methods that are declared but never used. Enabled at the level `warn`.

- `prefer-object-has-own` at the level `off`. This rule favors the use of `Object.hasOwn()` over `Object.prototype.hasOwnProperty.call()` a pattern we do not typically use. Furthermore, it is only relevant in an ES2020 context which some projects may not conform to. (added v8.5.0)
- Set the new `enforceForClassFields` property on `class-methods-use-this` to `false` to allow for arrow function class methods.
- Set the new `onlyOneSimpleParam` property on `no-confusing-arrow` to false. This explicitly chooses the current default behavior.
- Set the new `destructuredArrayIgnorePattern` property on `no-unused-vars` to `^_`. This allows us to indicate unused array elements when destructuring with a leading underscore.

## New to ESLint Plugin React

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
<MyComponent jsxProp={<div />} />; // ✅ Element passed as prop should be wrapped in curly braces.
<MyComponent jsxProp=<div /> />; // ❌ Valid JSX but confusing to read.
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

Finally, several existing plugin rules that had been unwittingly omitted in previous versions of this configuration we added.

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

- Enable the `require-atomic-updates` rule at the `warn` level. As we deal with more asynchronous operations these sorts of bugs become more of a concern.
- Use a stricter configuration of the `no-inner-declarations` rule disallowing variable declarations as well as function declarations. We typically use block bound `let` and `const` declarations so this may be a bit redundant, but it doesn't hurt.
- Use a stricter configuration of the `no-shadow` rule by setting the `builtinGlobals` property to `true`. This disables shadowing of global variables (ex. `Object` and `Array`).

The missing rules `import/no-relative-packages` and `import/no-import-module-exports` were added but at the level `off` so they will have no impact.
