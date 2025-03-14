Set-Location $PSScriptRoot

function deployUsingLocalAddons() {
    $azureSDKToolsRoot = "<Git clone of azure-sdk-tools>"
    $stressTestAddonsFolder = "$azureSDKToolsRoot/tools/stress-cluster/cluster/kubernetes/stress-test-addons"
    $clusterResourceGroup = "<Resource Group for Cluster>"
    $clusterSubscription = "<Azure Subscription>"
    $helmEnv = "pg2"

    if (-not (Get-ChildItem $stressTestAddonsFolder)) {
        Write-Host "Can't find the the new stress test addons folder at $stressTestAddonsFolder"
        return
    }

    pwsh "$azureSDKToolsRoot/eng/common/scripts/stress-testing/deploy-stress-tests.ps1" `
        -LocalAddonsPath "$stressTestAddonsFolder"  `
        -clusterGroup "$clusterResourceGroup" `
        -subscription "$clusterSubscription" `
        -Environment $helmEnv
}

#deployUsingLocalAddons
$gitRoot = git rev-parse --show-toplevel
pwsh "$gitRoot/eng/common/scripts/stress-testing/deploy-stress-tests.ps1" @args
