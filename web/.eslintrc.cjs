module.exports = {
    //"env": {
    //    "browser": true,
    //    "es2021": true
    //},
    "parser": "vue-eslint-parser",
    "extends": [
        "eslint:recommended",
        //"plugin:vue/essential",
        "plugin:@typescript-eslint/recommended",
        "plugin:vue/vue3-recommended",
    ],
    "parserOptions": {
        "ecmaVersion": 2020,
        "parser": "@typescript-eslint/parser",
        "sourceType": "module"
    },
    "plugins": [
        "vue",
        "@typescript-eslint"
    ],
    "rules": {
        "quotes": ["error", "double",
            {"avoidEscape": true, "allowTemplateLiterals": true}
        ],
        "semi": ["error", "never"],
        "vue/multi-word-component-names": "off"
    }
}
