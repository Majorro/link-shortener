import './App.css';
import React, {useState} from "react";

function App() {
  const [URL, setURL] = useState('')
  const [shortedURL, setShortedURL] = useState('')
  const [statusText, setStatusText] = useState('')
  const [isError, setIsError] = useState(false)

  const resetResults = () => {
    setIsError(false)
    setShortedURL('')
    setStatusText('')
  }

  const handleUrlSubmit = async e => {
    e.preventDefault()

    resetResults()

    if (!URL.length) {
      return
    }

    const APIBaseURL = process.env.REACT_APP_API_BASE_URL
    const {error, errorString, outputLink} = await (await fetch(APIBaseURL, {
      method: 'POST',
      body: JSON.stringify({
        link: URL
      })
    })).json()

    if (error) {
      setStatusText(errorString)
      setIsError(true)
    } else {
      setShortedURL(outputLink)
    }
  }

  const handleCopy = () => {
    navigator.clipboard.writeText(shortedURL)
    setStatusText('Copied!')
  }

  return (
    <div className="App">
      <header className="App-header">
        <div style={{"marginTop": 15}}>
          <h2>URL shortener</h2>
        </div>
        <form onSubmit={handleUrlSubmit}>
          <div style={{'display': 'flex', 'flexDirection': 'row'}}>
            <label>
              <div style={{'display': 'flex', 'flexDirection': 'row'}}>
                URL:
                <div style={{'width': '10px'}}/>
                <input type="url" value={URL} onChange={(e) => setURL(e.target.value)}/>
              </div>
            </label>
            <div style={{'width': '10px'}}/>
            <input type="submit" value="Shorten!"/>
          </div>
        </form>
        <div style={{'height': '10px'}}/>
        {shortedURL.length ?
          <>
            <div style={{'display': 'flex', 'flexDirection': 'row'}}>
              <a target='_blank' rel='noreferrer' href={shortedURL} className='App-link'>
                {shortedURL}
              </a>
              <div style={{'width': '10px'}}/>
              <button onClick={handleCopy}> Copy</button>
            </div>
            <div style={{'height': '10px'}}/>
          </>
          : <></>}
        {statusText.length ?
          <div>
            {`${isError ? 'Server error: ' : ''}${statusText}`}
          </div> : <></>
        }
      </header>
    </div>
  );
}

export default App;
