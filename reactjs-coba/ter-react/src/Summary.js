import Picme from './assets/Safiar.jpg';

function Summary() {
  return (
    <section id="summary" class="personal-info text-center">
      <figure class="flex justify-center">
        <img class="rounded-full sm:w-1/4 md:w-1/6" src={Picme} width="200px"
          alt="Foto Formal Saya"/>
      </figure>
      <h1 class="text-xl font-bold">Safiar Azalia Amanullah</h1>
      <p>An undergraduate student in Informatics Engineering, Dian Nuswantoro University
        Semarang. </p>
      <p> I
        have an interest in technologies such as Full-Stack Developer and Artificial Intelligence,
        therefore I want to learn a lot about both.</p>
      <p> I am a responsible person and eager to explore new
        things.</p>
    </section>
  );
}

export { Summary };