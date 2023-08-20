import React from "react";

function Product() {
    return (
        <div className="card w-full bg-base-100 shadow-xl rounded-none card-compact">
            <figure>
                <img src="/hbg.webp" alt="Shoes" />
            </figure>
            <div className="card-body">
                <h2 className="card-title">
                    Hamburder!
                    <div className="badge badge-primary">39$</div>
                </h2>
                <p>Grocery store </p>
                <div className="card-actions justify-end ">
                    <button className="btn btn-accent">Add to Cart</button>
                </div>
            </div>
        </div>
    );
}

export default Product;
