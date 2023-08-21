import React from 'react'
import Nav from './Nav'
import { NavLink } from 'react-router-dom'

const navLink = ["All", "Food", "Snack", "Drink", "Toy", "Vegetable"]

function Layout({ children }) {    
  return (
    <div className=''>
        <Nav/>
        <div className='flex justify-center p-14'>
            <div className='flex w-full rounded-md'>
                <div>
                <div className='w-96 border flex flex-col shadow-md ease-in bg-base-100'>
                    <NavLink className="navlnk" to="/">All</NavLink>
                    <NavLink className="navlnk" to="/food">Food 🌭</NavLink>
                    <NavLink className="navlnk" to="/snack">Snack 🍿</NavLink>
                    <NavLink className="navlnk" to="/drink">Drink 🧃</NavLink>
                    <NavLink className="navlnk" to="/toy">Toy 🧸</NavLink>
                    <NavLink className="navlnk" to="/vegetable">Vegetable 🥦</NavLink>
                    <NavLink className="navlnk" to="/cart">My CART 🛒</NavLink>
                </div>
                </div>
                <div className='w-full px-10'>
                    {children}
                </div>
            </div>
        </div>
    </div>
  )
}

export default Layout