export interface Handler {
    handle(msg: string): Promise<void>
}

