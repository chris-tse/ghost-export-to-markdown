const os = require('os')
const fs = require('fs')
const path = require('path')

try {
  process.stderr.write('Running install script...\n')
  const platform = os.platform()
  const binSuffix = platform === 'win32' ? 'windows.exe' : 
                    platform === 'darwin' ? 'macos' : 'linux'

  const sourcePath = path.join(__dirname, '..', 'bin', `cli-${binSuffix}`)
  const targetPath = path.join(__dirname, '..', 'bin', 'cli')

  process.stderr.write(`Installing for platform: ${platform}\n`)
  process.stderr.write(`Copying from: ${sourcePath}\n`)
  process.stderr.write(`Copying to: ${targetPath}\n`)

  if (!fs.existsSync(sourcePath)) {
    console.error(`Binary not found for platform: ${platform}`)
    process.exit(1)
  }

  fs.copyFileSync(sourcePath, targetPath)
  fs.chmodSync(targetPath, 0o755)
  
  process.stderr.write(`Successfully installed binary for ${platform}\n`)
} catch (error) {
  console.error('Installation failed:', error)
  process.exit(1)
}