import React from "react";

function Header() {
  return (
    <header className="normal flex flex-col top-0 font-serif">
      <h1 className="font-serif font-bold text-xl sm:text-5xl flex justify-center">
        Curriculum Vitae
      </h1>
      <nav className="flex flex-wrap flex-row justify-center">
        <table>
          <tr>
            <th>
              <a
                href="#summary"
                className="rounded-full p-1 hover:bg-black hover:text-white"
              >
                Summary
              </a>
            </th>
            <th>
              <a
                className="rounded-full p-1 hover:bg-black hover:text-white"
                href="#education"
              >
                Education
              </a>
            </th>
            <th>
              <a
                className="rounded-full p-1 hover:bg-black hover:text-white"
                href="#portofolio"
              >
                Portofolio
              </a>
            </th>
            <th>
              <a
                className="rounded-full p-1 hover:bg-black hover:text-white"
                href="#contact"
              >
                Contact
              </a>
            </th>
          </tr>
        </table>
      </nav>
    </header>
  );
}

function Footer() {
  return (
    <footer id="contact" className="flex flex-col items-center md:mx-auto mt-4">
      <table>
        <tr>
          <th>
            <a
              className="rounded-full p-1 hover:bg-black hover:text-white"
              target="_blank"
              rel="noopener noreferrer"
              href="https://github.com/safiarazl"
            >
              Github
            </a>
          </th>
          <th>
            <a
              className="rounded-full p-1 hover:bg-black hover:text-white"
              target="_blank"
              rel="noopener noreferrer"
              href="https://www.linkedin.com/mwlite/in/safiar-azalia-091828169"
            >
              LinkedIn
            </a>
          </th>
          <th>
            <a
              className="rounded-full p-1 hover:bg-black hover:text-white"
              target="_blank"
              rel="noopener noreferrer"
              href="https://www.instagram.com/fial_azl/"
            >
              Instagram
            </a>
          </th>
        </tr>
      </table>
      <p>© Safiar Azalia Amanullah - 2023</p>
    </footer>
  );
}

export { Header, Footer };
