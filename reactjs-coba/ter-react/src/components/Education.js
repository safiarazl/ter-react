import React from 'react';

const Organization = [
    {
        stage: "kuliah",
        name: "Universitas Dian Nuswantoro Semarang",
        jurusan: "Teknik Informatika",
        year: "2020 - Sekarang",
    },
    {
        stage: "sma",
        name: "SMA Negeri 2 Kudus",
        jurusan: "MIPA",
        year: "2017 - 2020",
    },
    {
        stage: "smp",
        name: "SMP Pondok Modern Selamat Kendal",
        jurusan: "",
        year: "2014 - 2017",
    }
]

function Education() {
    return (
<section id="education">
<h1 class="font-bold text-md sm:text-xl ">Latar belakang</h1>
<div class="pendidikan flex flex-col md:flex-row md:justify-around">
    {Organization.map((org) => 
        <div class={org.stage}>
        <h2 class="font-bold text-sm sm:text-lg md:text-xl">{org.name}</h2>
        <p class="text-xs sm:text-md md:text-lg">{org.jurusan === "" ? "SSN" : org.jurusan}</p>
        <p class="text-xs sm:text-md md:text-lg">{org.year}</p>
    </div>
    )}
</div>
</section>
    );
}

export { Education };