import { configureWunderGraphApplication, cors, EnvironmentVariable, introspect, templates } from '@wundergraph/sdk';
import server from './wundergraph.server';
import operations from './wundergraph.operations';

const federation = introspect.federation({
	upstreams: [
		{
			name: 'accounts',
			url: 'http://localhost:4001/graphql',
		},
		{
			name: 'reviews',
			url: 'http://localhost:4002/graphql',
		},
		{
			name: 'products',
			url: 'http://localhost:4003/graphql',
		},
		{
			name: 'inventory',
			url: 'http://localhost:4004/graphql',
		},
	],
});


// configureWunderGraph emits the configuration
configureWunderGraphApplication({
	apis: [federation],
	server,
	operations,
	codeGenerators: [
		{
			templates: [
				// use all the typescript react templates to generate a client
				...templates.typescript.all,
			],
		},
	],
	cors: {
		...cors.allowAll,
		allowedOrigins:
			process.env.NODE_ENV === 'production'
				? [
						// change this before deploying to production to the actual domain where you're deploying your app
						'http://localhost:3000',
				  ]
				: ['http://localhost:3000', new EnvironmentVariable('WG_ALLOWED_ORIGIN')],
	},
	security: {
		enableGraphQLEndpoint: true,
	},
});
