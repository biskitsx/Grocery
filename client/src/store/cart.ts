import { createSlice } from "@reduxjs/toolkit";

export interface CartDTO {
    ID: number
    categoryID: 1
    name: string
    picture: string
    price: string
    CreatedAt: Date
    UpdatedAt: Date
    DeletedAt: Date
}

const cartSlice = createSlice({
    name: "cart",
    initialState: {
        cart: [] as CartDTO[]
    },
    reducers: {
        load: (state, action) => {
            state.cart = action.payload
        },
        add: (state, action) => {
            state.cart = [action.payload, ...state.cart]
        },
        remove: (state, action) => {
            const cartTemp: CartDTO[] = []
            state.cart.forEach((item) => {
                if (action.payload.ID != item.ID) {
                    cartTemp.push(item)
                }
            })
            state.cart = cartTemp
        }
    }
})

export const { load, add, remove } = cartSlice.actions
export default cartSlice.reducer