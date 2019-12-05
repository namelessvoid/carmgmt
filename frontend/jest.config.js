module.exports = {
    transform: {
        '^.+.js$': 'babel-jest',
        '^.+.svelte$': 'jest-transform-svelte'
    },
    moduleFileExtensions: ['js','svelte'],
    "transformIgnorePatterns": [
        // Also transpile svelte-routing
        "node_modules/(?!(svelte-routing)/)"
    ]
};