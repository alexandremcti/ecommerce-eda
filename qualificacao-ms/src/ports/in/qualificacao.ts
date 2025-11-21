export type QualificacaoServicePortInputDTO = {
    messageId: string
    id: string
    totalAmount: number
    status: string
    items: Array<{
        productId: string
        count: number
        price: number
    }>
    customer:{
        id: string
        firstName: string
        lastName: string
        email: string
        deliveryAddress: {
            street: string
            number: number
            city: string
            state: string
            postalCode: string
        }
    }
    payment:{
        cardId: string
        bin: string
        number_token: string
        cardholder_name: string
        security_code: string
        expiration_month: string
        expiration_year: string
        brand: string
    }
    createdAt: Date
    updatedAt: Date
}

export interface QualificacaoServicePort {
    create(input: QualificacaoServicePortInputDTO): Promise<void>
}