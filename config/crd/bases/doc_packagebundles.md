# API Reference

Packages:

- [packages.eks.amazonaws.com/v1alpha1](#packageseksamazonawscomv1alpha1)

# packages.eks.amazonaws.com/v1alpha1

Resource Types:

- [PackageBundle](#packagebundle)




## PackageBundle
<sup><sup>[↩ Parent](#packageseksamazonawscomv1alpha1 )</sup></sup>






PackageBundle is the Schema for the packagebundles API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>packages.eks.amazonaws.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PackageBundle</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlespec">spec</a></b></td>
        <td>object</td>
        <td>
          PackageBundleSpec defines the desired state of PackageBundle.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#packagebundlestatus">status</a></b></td>
        <td>object</td>
        <td>
          PackageBundleStatus defines the observed state of PackageBundle.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.spec
<sup><sup>[↩ Parent](#packagebundle)</sup></sup>



PackageBundleSpec defines the desired state of PackageBundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#packagebundlespecpackagesindex">packages</a></b></td>
        <td>[]object</td>
        <td>
          Packages supported by this bundle.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PackageBundle.spec.packages[index]
<sup><sup>[↩ Parent](#packagebundlespec)</sup></sup>



BundlePackage specifies a package within a bundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#packagebundlespecpackagesindexsource">source</a></b></td>
        <td>object</td>
        <td>
          Source location for the package (probably a helm chart).<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the package.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.spec.packages[index].source
<sup><sup>[↩ Parent](#packagebundlespecpackagesindex)</sup></sup>



Source location for the package (probably a helm chart).

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository within the Registry where the package is found.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlespecpackagesindexsourceversionsindex">versions</a></b></td>
        <td>[]object</td>
        <td>
          Versions of the package supported by this bundle.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>registry</b></td>
        <td>string</td>
        <td>
          Registry in which the package is found.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.spec.packages[index].source.versions[index]
<sup><sup>[↩ Parent](#packagebundlespecpackagesindexsource)</sup></sup>



SourceVersion describes a version of a package within a repository.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>digest</b></td>
        <td>string</td>
        <td>
          Digest is a checksum value identifying the version of the package and its contents.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is a human-friendly description of the version, e.g. "v1.0".<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlespecpackagesindexsourceversionsindeximagesindex">images</a></b></td>
        <td>[]object</td>
        <td>
          Images is a list of images used by this version of the package.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>schema</b></td>
        <td>string</td>
        <td>
          Schema is a base64 encoded, gzipped json schema used to validate package configurations.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.spec.packages[index].source.versions[index].images[index]
<sup><sup>[↩ Parent](#packagebundlespecpackagesindexsourceversionsindex)</sup></sup>



VersionImages is an image used by a version of a package.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>digest</b></td>
        <td>string</td>
        <td>
          Digest is a checksum value identifying the version of the package and its contents.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository within the Registry where the package is found.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PackageBundle.status
<sup><sup>[↩ Parent](#packagebundle)</sup></sup>



PackageBundleStatus defines the observed state of PackageBundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>state</b></td>
        <td>enum</td>
        <td>
          PackageBundleStateEnum defines the observed state of PackageBundle.<br/>
          <br/>
            <i>Enum</i>: available, ignored, invalid<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlestatusspec">spec</a></b></td>
        <td>object</td>
        <td>
          PackageBundleSpec defines the desired state of PackageBundle.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.status.spec
<sup><sup>[↩ Parent](#packagebundlestatus)</sup></sup>



PackageBundleSpec defines the desired state of PackageBundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#packagebundlestatusspecpackagesindex">packages</a></b></td>
        <td>[]object</td>
        <td>
          Packages supported by this bundle.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PackageBundle.status.spec.packages[index]
<sup><sup>[↩ Parent](#packagebundlestatusspec)</sup></sup>



BundlePackage specifies a package within a bundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#packagebundlestatusspecpackagesindexsource">source</a></b></td>
        <td>object</td>
        <td>
          Source location for the package (probably a helm chart).<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name of the package.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.status.spec.packages[index].source
<sup><sup>[↩ Parent](#packagebundlestatusspecpackagesindex)</sup></sup>



Source location for the package (probably a helm chart).

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository within the Registry where the package is found.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlestatusspecpackagesindexsourceversionsindex">versions</a></b></td>
        <td>[]object</td>
        <td>
          Versions of the package supported by this bundle.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>registry</b></td>
        <td>string</td>
        <td>
          Registry in which the package is found.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.status.spec.packages[index].source.versions[index]
<sup><sup>[↩ Parent](#packagebundlestatusspecpackagesindexsource)</sup></sup>



SourceVersion describes a version of a package within a repository.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>digest</b></td>
        <td>string</td>
        <td>
          Digest is a checksum value identifying the version of the package and its contents.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>name</b></td>
        <td>string</td>
        <td>
          Name is a human-friendly description of the version, e.g. "v1.0".<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlestatusspecpackagesindexsourceversionsindeximagesindex">images</a></b></td>
        <td>[]object</td>
        <td>
          Images is a list of images used by this version of the package.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>schema</b></td>
        <td>string</td>
        <td>
          Schema is a base64 encoded, gzipped json schema used to validate package configurations.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundle.status.spec.packages[index].source.versions[index].images[index]
<sup><sup>[↩ Parent](#packagebundlestatusspecpackagesindexsourceversionsindex)</sup></sup>



VersionImages is an image used by a version of a package.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>digest</b></td>
        <td>string</td>
        <td>
          Digest is a checksum value identifying the version of the package and its contents.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository within the Registry where the package is found.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>