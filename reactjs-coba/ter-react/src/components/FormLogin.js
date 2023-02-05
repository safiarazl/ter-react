import React from "react";
import { Footer } from "./Heafoot";

function Form({username, password, handleChange, handleSubmit}) {
    return (
        <main className="bg-white m-20 rounded-xl flex flex-col p-10 shadow-2xl shadow-black w-96">
            <h1  className="block text-center font-serif pt-3 h-16 text-4xl font-bold">
            <i className="fa-solid fa-user"></i> Login
            </h1>
            <form onSubmit={handleSubmit}>
            <div className="flex flex-col">
                <div className="username">
                    <label>Username :</label>
                    <input
                        className="mt-1 block w-full px-3 py-2 bg-white border border-slate-300 rounded-md text-sm shadow-sm"
                        type="text" placeholder="Type username here" name="username" value={ username } onChange={(event) => handleChange(event)} />
                </div>
                <div className="password">
                    <label>Password :</label>
                    <input
                        className="mt-1 block w-full px-3 py-2 bg-white border border-slate-300 rounded-md text-sm shadow-sm"
                        type="password" placeholder="Type password here" name="password" value={password} onChange={(event) => handleChange(event) } />
                </div>
            </div>
            <div className="flex justify-center pt-1 mt-3">
                <button className=" rounded-full border-2 w-full h-full hover:bg-black hover:text-white" type="submmit">Login</button>
            </div>
        </form>
        <Footer />
        </main>
    );
}

export default Form;