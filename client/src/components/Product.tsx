import axios from "axios";
import React, { SyntheticEvent } from "react";
import { useDispatch, useSelector } from "react-redux";
import { add, remove } from "../store/cart";

function Product({ props }) {
    const dispatch = useDispatch()
    const cart = useSelector((state) => state.cart)
    const handleButton = async (e: SyntheticEvent) => {
        e.preventDefault()
        const res = await axios.put(`http://localhost:3001/api/product/add-remove/${props.ID}`)
        const removeCart = res.data.remove
        if (removeCart) {
            dispatch(remove(props))
        } else {
            dispatch(add(props))
        }
        console.log(cart)
    }

    return (
        <div className="card w-full bg-base-100 shadow-xl rounded-none card-compact">
            <figure className="h-52">
                <img src={props.picture} alt="Shoes" className="h-full w-full" />
            </figure>
            <div className="card-body">
                <h2 className="card-title">
                    {props.name}!
                    <div className="badge badge-primary">{props.price}$</div>
                </h2>
                <p>Grocery store </p>
                <div className="card-actions justify-end ">
                    <button className="btn btn-accent" onClick={handleButton}>Add to Cart</button>
                </div>
            </div>
        </div>
    );
}

export default Product;
