const os = require('os')
const fs = require('fs')

const platform = os.platform()
const binSuffix = platform === 'win32' ? 'windows.exe' : 
                  platform === 'darwin' ? 'macos' : 'linux'

fs.copyFileSync(
  `bin/cli-${binSuffix}`,
  'bin/cli'
)
fs.chmodSync('bin/cli', 0x755)