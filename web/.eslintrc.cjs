module.exports = {
	'env': {
		'browser': true,
		'es2021': true,
		'node': true
	},
	'extends': [
		'eslint:recommended',
		'plugin:vue/vue3-essential'
	],
	'parserOptions': {
		'ecmaVersion': 12,
		'sourceType': 'module'
	},
	'plugins': [
		'vue'
	],
	'rules': {
        "quotes": ["error", "double",
            {"avoidEscape": true, "allowTemplateLiterals": true}
        ],
		"semi": ["error", "never"],
        "vue/multi-word-component-names": "off",
		"indent": ["error", 2],
		"vue/html-indent": ["error", 2]
	}
}