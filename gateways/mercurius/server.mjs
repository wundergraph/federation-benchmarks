import Fastify from 'fastify'
import mercuriusWithGateway from '@mercuriusjs/gateway'

const gateway = Fastify()

gateway.register(mercuriusWithGateway, {
    gateway: {
        services: [
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
        pollingInterval: 2000
    },
})

gateway.listen({ port: 3013 })
