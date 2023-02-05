import React from 'react';

const Portofolios = [
    {
        id: 1,
        name: "Portofolio name",
        description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Perferendis illum obcaecati nostrum atque, dolorem incidunt explicabo quam quae maxime placeat!",
        image: "https://dummyimage.com/400x300/000/fff.png&text=Foto+Portofolio",
    },
]

function loopPortofolio(num) {
    const card = Portofolios.map((porto) =>
                    <div class="border-solid rounded-2xl shadow-md shadow-gray-600">
                        <img class="rounded-t-lg" src={porto.image} alt={`image porto ${porto.name}`}/>
                        <div class="padding p-2">
                            <h2 class="font-bold">{porto.name}</h2>
                            <p>{porto.description}</p>
                        </div>
                    </div>
                )
    let rows = [];
    for (let i = 0; i < num; i++) {
        rows.push(card);
    }
    return rows;
}

function Portofolio() {
    return (
        <section id="portofolio">
            <h1 class="font-bold text-md sm:text-xl ">Portofolio</h1>
            <div class="container-porto portofolio grid grid-rows-1 gap-4 sm:grid-cols-2 md:grid-cols-3">
                { loopPortofolio(999) }
            </div>
        </section>
    );
}

export { Portofolio };