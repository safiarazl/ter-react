import React from 'react';
import { Summary } from "../components/Summary";
import { Education } from "../components/Education";
import { Portofolio } from "../components/Porto";
import { Header, Footer } from "../components/Heafoot";

function Contents() {
  return (
  <main className="max-w-sm sm:max-w-md md:max-w-5xl mx-auto">
    <Summary/>
    <Education/>
    <Portofolio/>
  </main>
  );
}

function Cv() {
    return (
        <div className="bg-white m-10 rounded-xl p-5 shadow-2xl shadow-black">
            <Header />
            <Contents />
            <Footer />
        </div>
      );
}

export default Cv;