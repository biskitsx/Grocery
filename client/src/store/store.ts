import { configureStore } from "@reduxjs/toolkit";
import userSlice from './user'
import cartSlice from './cart'
export default configureStore({
    reducer: {
        user: userSlice,
        cart: cartSlice
    }
})