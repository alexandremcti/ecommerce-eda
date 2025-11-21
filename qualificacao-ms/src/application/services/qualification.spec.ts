import { faker } from '@faker-js/faker/locale/pt_BR'
import { QualificacaoServicePortInputDTO } from '../../ports/in/qualificacao'
import { OrderDTO, QualificacaoOutput } from '../../ports/out/qualificacao'
import { QualificacaoService } from './qualification'

describe("QualificationService", () => {
    it("Should update order status to qualified", async () => {
        const {sut, qualificacaoOutput} = createSut()
        const input = makeInputDTO()
        await sut.create(input)
        const output = qualificacaoOutput.orderEvents.find(q => q.data.id === input.id)

        expect(output?.data.id).toBeDefined()
        expect(output?.data.status).toBe('qualificado')
    })


        it("Should update order status to not qualified", async () => {
        const {sut, qualificacaoOutput} = createSut()
        const input = makeInputDTO()
        input.customer.deliveryAddress.state = 'ZZ'
        await sut.create(input)
        const output = qualificacaoOutput.orderEvents.find(q => q.data.id === input.id)

        expect(output?.data.id).toBeDefined()
        expect(output?.data.status).toBe('recusado')
    })
});




type Event = {
    eventName: string
    data: OrderDTO
}

class QualificacaoOutputMock implements QualificacaoOutput {
    orderEvents: Event[] = []

    async notify(data: Event): Promise<void> {
        this.orderEvents.push(data)
    }

    async OrderQualified(data: OrderDTO): Promise<void> {
        this.notify({eventName: 'OrderQualified', data})
    }
    
    async OrderRecused(data: OrderDTO): Promise<void> {
        this.notify({eventName: 'OrderRecused', data})
    }

}

type Sut = {
    sut: QualificacaoService
    qualificacaoOutput: QualificacaoOutputMock
}

const createSut = (): Sut => {
    const qualificacaoOutput = new QualificacaoOutputMock()
    const sut = new QualificacaoService(qualificacaoOutput)
    return {
        sut,
        qualificacaoOutput
    }
}

const makeInputDTO = (): QualificacaoServicePortInputDTO => {
    return {
        messageId: faker.string.uuid(),
        id: faker.string.uuid(),
        totalAmount: faker.number.float({min: 10, max: 1000}),
        status: 'recebido',
        items: [
            {
                productId: faker.string.uuid(),
                count: faker.number.int({min: 1, max: 10}),
                price: faker.number.float({min: 10, max: 500})
            }
        ],
        customer: {
            id: faker.string.uuid(),
            firstName: faker.person.firstName(),
            lastName: faker.person.lastName(),
            email: faker.internet.email(),
            deliveryAddress: {
                street: faker.location.street(),
                number: faker.number.int({min: 1, max: 1000}),
                city: faker.location.city(),
                state: faker.location.state({abbreviated: true }),
                postalCode: faker.location.zipCode()
            }
        },
        payment: {
            cardId: faker.string.uuid(),
            bin: faker.finance.creditCardNumber(),
            number_token: faker.string.uuid(),
            cardholder_name: faker.person.fullName(),
            security_code: faker.string.numeric(3),
            expiration_month: faker.date.month({ abbreviated: true }),
            expiration_year: faker.date.future().getFullYear().toString(),
            brand: faker.vehicle.manufacturer()
        },
        createdAt: new Date(),
        updatedAt: new Date()
    }
}
