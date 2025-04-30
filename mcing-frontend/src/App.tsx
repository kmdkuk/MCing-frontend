import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const [pods, setPods] = useState<any[]>([]);
  const [minecrafts, setMinecrafts] = useState<any[]>([]);

  const fetchPods = async () => {
    try {
      const response = await fetch('http://localhost:8081/pods');
      const data = await response.json();
      setPods(data.items || []);
    } catch (error) {
      console.error('Error fetching pods:', error);
    }
  };

  const fetchMinecrafts = async () => {
    try {
      const response = await fetch('http://localhost:8081/minecrafts');
      const data = await response.json();
      setMinecrafts(data.items || []);
      console.log(data);
    } catch (error) {
      console.error('Error fetching minecrafts:', error);
    }
  };
  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <div>
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <button onClick={fetchPods}>Fetch Pods</button>
          <button onClick={fetchMinecrafts}>Fetch Minecrafts</button>
        </div>

        <div>
          <h2>Pods</h2>
          <ul>
            {pods.map((pod) => (
              <li key={pod.metadata.name}>{pod.metadata.namespace + "/" + pod.metadata.name}</li>
            ))}
          </ul>
        </div>

        <div>
          <h2>Minecrafts</h2>
          <ul>
            {minecrafts.map((minecraft) => (
              <li key={minecraft.metadata.name}>
                {minecraft.metadata.namespace + "/" + minecraft.metadata.name}
              </li>
            ))}
          </ul>
        </div>

        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
