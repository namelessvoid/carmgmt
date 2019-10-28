module.exports = {
    transform: {
        '^.+.js$': 'babel-jest',
        '^.+.svelte$': 'jest-transform-svelte'
    },
    moduleFileExtensions: ['js','svelte'],
    "transformIgnorePatterns": [
        "node_modules/(?!(svelte-routing)/)"
    ]
};