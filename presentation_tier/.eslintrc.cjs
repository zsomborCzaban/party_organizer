module.exports = {
    root: true,
    parserOptions: {
      sourceType: 'module',
      project: './tsconfig.eslint.json',
    },
    extends: ['airbnb-base', 'airbnb-typescript/base', 'prettier', 'plugin:react/recommended', 'plugin:react/jsx-runtime'],
    plugins: ['simple-import-sort', 'check-file', 'prefer-arrow-functions', 'react-hooks', 'path'],
    rules: {
      'quotes': ['error', 'single'],
      'class-methods-use-this': 'warn',
      'import/prefer-default-export': 'off',
      'no-param-reassign': 'off',
      'linebreak-style': 'off',
      'max-len': ['warn', { code: 240 }],
      'no-plusplus': ['error', { allowForLoopAfterthoughts: true }],
      'jsx-quotes': ['error', 'prefer-single'],
      '@typescript-eslint/comma-dangle': ['error', 'always-multiline'],
      '@typescript-eslint/no-explicit-any': ['error'],
      '@typescript-eslint/no-unused-vars': ['error', { argsIgnorePattern: '^_' }],
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/switch-exhaustiveness-check': ['error'],
      'react/jsx-closing-bracket-location': ['error'],
      'react/jsx-curly-brace-presence': ['error', { props: 'never', children: 'ignore' }],
      'react/jsx-first-prop-new-line': ['error', 'multiline-multiprop'],
      'react/jsx-max-props-per-line': ['error', { when: 'multiline' }],
      'react/jsx-tag-spacing': [
        'error',
        {
          closingSlash: 'never',
          beforeSelfClosing: 'always',
          afterOpening: 'never',
          beforeClosing: 'allow',
        },
      ],
      'react/jsx-wrap-multilines': [
        'error',
        {
          arrow: 'parens-new-line',
        },
      ],
      'react/jsx-indent': ['error', 2],
      'react/jsx-indent-props': ['error', 2],
      'react/self-closing-comp': [
        'error',
        {
          component: true,
          html: true,
        },
      ],
      'semi': ['error'],
      'import/order': 'off',
      'prefer-arrow-functions/prefer-arrow-functions': [
        'error',
        {
          disallowPrototype: false,
          returnStyle: 'implicit',
          singleReturnOnly: false,
        },
      ],
      'react-hooks/rules-of-hooks': ['error'],
      'react-hooks/exhaustive-deps': ['error'],
      'path/no-relative-imports': 'error',
      'no-console': 'off',
      'no-restricted-syntax': [
        'error',
        // Enabling ForOfStatement, but keeping the rest of the airbnb defaults.
        {
          selector: 'ForInStatement',
          message:
            'for..in loops iterate over the entire prototype chain, which is virtually never what you want. Use Object.{keys,values,entries}, and iterate over the resulting array.',
        },
        {
          selector: 'LabeledStatement',
          message: 'Labels are a form of GOTO; using them makes code confusing and hard to maintain and understand.',
        },
        {
          selector: 'WithStatement',
          message: '`with` is disallowed in strict mode because it makes code impossible to predict and optimize.',
        },
      ],
      'no-nested-ternary': 'off',
    },
    settings: {
      'import/resolver': {
        typescript: {},
      },
      'react': {
        version: 'detect',
      },
    },
  };
  