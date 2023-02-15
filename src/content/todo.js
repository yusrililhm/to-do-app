import '../assets/css/index.css'

function Todo() {
    return (
        <section className='font-Poppins'>
            <form method='POST' action='/Add'>
                <input type='text' name='todo' id='todo' placeholder='Type your plan here' />
            </form>
        </section>
    );
}

export default Todo;