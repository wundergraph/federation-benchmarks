import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from '@apollo/server/standalone';
import { ApolloGateway, IntrospectAndCompose } from "@apollo/gateway";

// import {serializeQueryPlan} from '@apollo/query-planner';

const gateway = new ApolloGateway({
    supergraphSdl: new IntrospectAndCompose({
        // subgraphs: [
        //     {
        //         name: "accounts",
        //         url: 'https://wg-federation-demo-accounts.fly.dev/graphql',
        //     },
        //     {
        //         name: "products",
        //         url: 'https://wg-federation-demo-products.fly.dev/graphql',
        //     },
        //     {
        //         name: "reviews",
        //         url: 'https://wg-federation-demo-reviews.fly.dev/graphql',
        //     },
        //     {
        //         name: "inventory",
        //         url: 'https://wg-federation-demo-inventory.fly.dev/graphql',
        //     },
        // ],
        subgraphs: [
            {
                name: "accounts",
                url: 'http://localhost:4001/graphql',
            },
            {
                name: "products",
                url: 'http://localhost:4003/graphql',
            },
            {
                name: "reviews",
                url: 'http://localhost:4002/graphql',
            },
            {
                name: "inventory",
                url: 'http://localhost:4004/graphql',
            },
        ],
    }),
    // __exposeQueryPlanExperimental: true,
    // experimental_didResolveQueryPlan: function(options) {
    //     if (options.requestContext.operationName !== 'IntrospectionQuery') {
    //         console.log(serializeQueryPlan(options.queryPlan));
    //     }
    // }
});

const server = new ApolloServer({ gateway });

// Note the top-level await!
const { url } = await startStandaloneServer(server, {
    listen: {
        port: 3010,
    }
});
console.log(`🚀  Server ready at ${url}`);
