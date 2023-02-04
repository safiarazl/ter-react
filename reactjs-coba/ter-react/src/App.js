import { Summary } from "./Summary";
import { Education } from "./Education";
import { Portofolio } from "./Porto";


function Header() {
  return (
    <header class="normal flex flex-col top-0 font-serif">
            <h1 class="font-serif font-bold text-xl sm:text-5xl flex justify-center">Curriculum Vitae</h1>
            <nav class="flex flex-wrap flex-row justify-center">
                <tr>
                    <th><a href="#summary" class="rounded-full p-1 hover:bg-black hover:text-white">Summary</a></th>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white" href="#education">Education</a></th>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white" href="#portofolio">Portofolio</a></th>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white"
                            href="#contact">Contact</a></th>
                </tr>
            </nav>
        </header>
  );
}

function Contents() {
  return (
  <main class="max-w-sm sm:max-w-md md:max-w-5xl mx-auto">
    <Summary/>
    <Education/>
    <Portofolio/>
  </main>
  );
}

function Footer() {
  return (
    <footer id="contact" class="flex flex-col items-center md:mx-auto mt-4">
            <table>
                <tr>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white" target="_blank"
                            href="https://github.com/safiarazl">Github</a></th>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white" target="_blank"
                            href="https://www.linkedin.com/mwlite/in/safiar-azalia-091828169">LinkedIn</a>
                    </th>
                    <th><a class="rounded-full p-1 hover:bg-black hover:text-white" target="_blank"
                            href="https://www.instagram.com/fial_azl/">Instagram</a></th>
                </tr>
            </table>
            <p>© Safiar Azalia Amanullah - BE2671558 - 2022</p>
        </footer>
  );
}

function App() {
  return (
    <div className="App" class="bg-white m-10 rounded-xl p-5 shadow-2xl shadow-black">
      <Header />
      <Contents />
      <Footer />
    </div>
  );
}

export default App;
