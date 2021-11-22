---
title: Document Data Structure
date: 2021-11-02
excerpt: 'In preparation of adding Commons support for reports, we review the existing file data structures to inform the creation of a new Document type.'
---

## Summary

In this study we look at potential changes to the data structure to rationalize/synchronize the way we store data on given content types. The origin of this investigation is a new requirement from the Research and Analytics Team that reports (a new document type) can have multiple teams as their authors. Currently document files do not have an team property, so we could simply add an array of teams to the data model. However, the team used this request as an opportunity to step back and take some time to evaluate the different ways we treat different content types with the intention of standardizing them a bit more

The general structure we would like to establish is a hierarchy based on the content specificity and complexity. The bulk of the content is composed of individual files. These files are grouped together in logic collection of varying degrees . This structure can be visualized with the pyramid seen below.

<img
alt="An isosceles triangle broken up into three horizontal sections. The top (smallest) triangle is labelled project level and described as a collection units and/or files. To the left of this triangle is a list of associated types - Graphic Project, Package, Playbook, and Video Project. The middle section is labelled unit level and described as the intermediate wrapper that adds additional metadata and can stand alone. To the left of this section is a list of associated types - Video Unit and Document with a question mark. The bottom (largest) section is labelled file level and described as the most atomic level which does not stand alone. To the left of this section is a list of associated types - Document File, Image File, Support File, and Video File."
src="{{ '/assets/2021/11/content-pyramid.svg' | relative_url }}"
style="max-width: 600px; text-align: center"
title="Infographic illustrating the content pyramid."
/>

## Top Level - Project MetaData

The highest (i.e. most complex) level of content is the project. Projects are composed of a collection of lower level data nodes and additional metadata. They represent collections of data curated in some meaningful fashion.

We currently have four project types, namely `GraphicProject`, `Package`, `Playbook`, and `VideoProject`. The below table demonstrates the metadata applied to each one of these project types. As you can see, there is at least on array of file or intermediate content types associated with each of these projects.

| Property           | GraphicProject     | Package             | Playbook           | VideoProject       |
| ------------------ | ------------------ | ------------------- | ------------------ | ------------------ |
| id                 | ✅                 | ✅                  | ✅                 | ✅                 |
| createdAt          | ✅                 | ✅                  | ✅                 | ✅                 |
| updatedAt          | ✅                 | ✅                  | ✅                 | ✅                 |
| publishedAt        | ✅                 | ✅                  | ✅                 | ✅                 |
| initialPublishedAt | ❌                 | ❌                  | ✅                 | ❌                 |
| alt                | ✅                 | ❌                  | ❌                 | ❌                 |
| assetPath          | ✅                 | ✅                  | ✅                 | ✅                 |
| author             | ✅ - User          | ✅ - User           | ✅ - User          | ✅ - User          |
| categories         | ✅                 | ✅                  | ✅                 | ✅                 |
| commonsResources   | ❌                 | ❌                  | ✅                 | ❌                 |
| content            | ❌                 | ❌                  | ✅                 | ❌                 |
| copyright          | ✅                 | ❌                  | ❌                 | ❌                 |
| countries          | ❌                 | ❌                  | ✅                 | ❌                 |
| desc               | ❌                 | ✅                  | ✅                 | ❌                 |
| descInternal       | ✅                 | ❌                  | ❌                 | ✅                 |
| descPublic         | ✅                 | ❌                  | ❌                 | ✅                 |
| documents          | ❌                 | ✅ - [DocumentFile] | ❌                 | ❌                 |
| images             | ✅ - [ImageFile]   | ❌                  | ❌                 | ❌                 |
| permalink          | ❌                 | ❌                  | ✅                 | ❌                 |
| policy             | ❌                 | ❌                  | ✅                 | ❌                 |
| projectTitle       | ❌                 | ❌                  | ❌                 | ✅                 |
| projectType        | ❌                 | ❌                  | ❌                 | ✅                 |
| socialResources    | ❌                 | ❌                  | ✅                 | ❌                 |
| status             | ✅                 | ✅                  | ✅                 | ✅                 |
| supportFiles       | ✅ - [SupportFile] | ❌                  | ✅ - [SupportFile] | ✅ - [SupportFile] |
| tags               | ✅                 | ✅                  | ✅                 | ✅                 |
| team               | ✅ - Team          | ✅ - Team           | ✅ - Team          | ✅ - Team          |
| thumbnails         | ❌                 | ❌                  | ❌                 | ✅ - [ImageFile]   |
| title              | ✅                 | ✅                  | ✅                 | ❌                 |
| type               | ✅                 | ✅                  | ✅                 | ❌                 |
| units              | ❌                 | ❌                  | ❌                 | ✅ - [VideoUnit]   |
| visibility         | ✅                 | ✅                  | ✅                 | ✅                 |

## Intermediary Level - Wrappers

An intermediary level wrapper is the bridge between a file and a project. It wraps a particular file (or group of files) in additional metadata. Items at this level can stand alone as a piece of content or can be combined to compose a project. Currently we only have an intermediary level for videos, namely the `VideoUnit` type. This video units contain the following properties:

- id
- createdAt
- updatedAt
- categories
- descPublic
- files - [VideoFile]
- language
- tags
- thumbnails - [Thumbnail]
- title

The `Thumbnail` type is also technically a intermediate type since it wraps an `ImageFile` with additional data pertaining to the image size, but it is very limited in it's scope.

## Bottom Level - File MetaData

The base of the pyramid (and the foundation for most content) is made up of specific files. These files should contain minimal metadata pertinent to just that specific file. Broader metadata (such as categories and tags) should be reserved for the project and intermediate level.

Unfortunately, we have been treating file-level types in different ways. The `DocumentFile` type in particular contains a lot of metadata that should be push up a level. This is was done so that document files could stand alone as a content type. In retrospect this was probably a mistake and we should rectify it by using a new document type where we have used document files in the past.

In the table below a 👍 indicates that this is an appropriate file level property, a 👎 indicates an inappropriate property, and 🤔 indicates an ambiguous case.

| Property               | DocumentFile     | ImageFile         | SupportFile         | VideoFile         |
| ---------------------- | ---------------- | ----------------- | ------------------- | ----------------- |
| id 👍                  | ✅               | ✅                | ✅                  | ✅                |
| createdAt 👍           | ✅               | ✅                | ✅                  | ✅                |
| updatedAt 👍           | ✅               | ✅                | ✅                  | ✅                |
| publishedAt 👎         | ✅               | ❌                | ❌                  | ❌                |
| alt 👍                 | ❌               | ✅                | ❌                  | ❌                |
| bitrate 👍             | ❌               | ❌                | ❌                  | ✅                |
| bureaus 👎             | ✅               | ❌                | ❌                  | ❌                |
| caption 👍             | ❌               | ✅                | ❌                  | ❌                |
| categories 👎          | ✅               | ❌                | ❌                  | ❌                |
| content 🤔             | ✅               | ❌                | ❌                  | ❌                |
| countries 👎           | ✅               | ❌                | ❌                  | ❌                |
| dimensions 👍          | ❌               | ✅                | ❌                  | ✅                |
| duration 👍            | ❌               | ❌                | ❌                  | ✅                |
| editable 👍            | ❌               | ❌                | ✅                  | ❌                |
| excerpt 🤔             | ✅               | ❌                | ❌                  | ❌                |
| filename 👍            | ✅               | ✅                | ✅                  | ✅                |
| filesize 👍            | ✅               | ✅                | ✅                  | ✅                |
| filetype 👍            | ✅               | ✅                | ✅                  | ✅                |
| image 👍               | ✅ - [ImageFile] | ❌                | ❌                  | ❌                |
| language 👍            | ✅               | ✅                | ✅                  | ✅                |
| longdesc 👍            | ❌               | ✅                | ❌                  | ❌                |
| md5 👍                 | ❌               | ✅                | ✅                  | ✅                |
| quality 👍             | ❌               | ✅                | ❌                  | ✅ - VideoQuality |
| signedUrl 👍           | ✅               | ✅                | ✅                  | ✅                |
| social 🤔              | ❌               | ✅                | ❌                  | ❌                |
| status 🤔              | ✅               | ❌                | ❌                  | ❌                |
| stream 👍              | ❌               | ❌                | ❌                  | ✅ - VideoStream  |
| style 🤔               | ❌               | ✅ - GraphicStyle | ❌                  | ❌                |
| tags 👎                | ✅               | ❌                | ❌                  | ❌                |
| title 🤔               | ✅               | ✅                | ❌                  | ❌                |
| url 👍                 | ✅               | ✅                | ✅                  | ✅                |
| use 🤔                 | ✅ - DocumentUse | ✅ - ImageUse     | ✅ - SupportFileUse | ✅ - VideoUse     |
| visibility 👍          | ✅               | ✅                | ✅                  | ✅                |
| videoBurnedInStatus 👍 | ❌               | ❌                | ❌                  | ✅                |

As can be seen in the above table, the `SupportFile` and `VideoFile` types are properly configured with only file specific data properties. The `DocumentFile` have multiple properties (bureaus, categories, countries, tags) that should be moved to an higher level type.

If Prisma 2 allows use to make use of extensible types we should create a common `File` type. This type would include properties common to all files and then be extended by `DocumentFile`, `ImageFile`, `SupportFile`, and `VideoFile`. The common properties are:

- id
- createdAt
- updatedAt
- filename
- filesize
- filetype
- language
- signedUrl
- url
- use (? - see the below [note on use](#uses))
- visibility

## The Document Type

In order to conform to the structure outlined above, we should create a new intermediate `Document` type. This type will accept an array of document files. It will also accept multiple teams and a broad range of tags. All told, the structure should look something like so:

```graphql
type Document {
  id: ID! @id
  createdAt: DateTime! @createdAt
  updatedAt: DateTime! @updatedAt
  publishedAt: DateTime
  title: String
  files: [DocumentFile] @relation(onDelete: CASCADE)
  bureaus: [Bureau]
  missions: [DiplomaticPost]
  categories: [Category]
  countries: [Country]
  tags: [Tag]
  policy: [PolicyPriority]
  team: [Team]
}
```

For the moment some of the more ambiguous properties (content, excerpt, status, and use) have been left off in favor of keeping them on the `DocumentFile`. However, we may want to revisit that decision as the publication flow for report files solidifies.

The use case for having multiple `DocumentFile` items in a single `Document` might included the same file contents in different formats (ex. docx and pdf) or possibly the same file in different languages.

In most cases we'll then want to migrate all arrays of `DocumentFile` to be arrays of `Document` with each `Document` associated with a single file. This migration must happen if we are to remove properties (ie. taxonomy items like categories and tags) from the `DocumentFile` type.

## Uses

Most of the file-specific types have a use property on them. This property indicates what category of document the file represents. The use types (`DocumentUse`, `GraphicStyle`, `ImageUse`, `SocialPlatform`, `SupportFileUse`, and `VideoUse`) are all simply a list of strings specifying the use name. Below is the list of existing uses.

<table>
  <caption><strong>Current Uses</strong></caption>
  <tr>
    <th>DocumentUse</th>
    <th>ImageUse</th>
    <th>SupportFileUse</th>
    <th>VideoUse</th>
    <th>ReportUse (Notional)</th>
  </tr>
  <tr>
    <td>
      Background Briefing<br>
      Department Press Briefing<br>
      Fact Sheet<br>
      Interview<br>
      Media Note<br>
      Notice to the Press<br>
      On-the-record Briefing<br>
      Press Guidance<br>
      Remarks<br>
      Speeches<br>
      Statement<br>
      Taken Questions<br>
      Travel Alert<br>
      Readout<br>
      Travel Warning
    </td>
    <td>
      Email Graphic<br>
      Infographic<br>
      Website Hero Image<br>
      Social Media Graphic<br>
      Memes<br>
      3D Graphics<br>
      Thumbnail/Cover Image
    </td>
    <td>
      N/A
    </td>
    <td>
      Clean<br>
      Full Video<br>
      Promotional Teaser<br>
      Video Assets (B-Roll)<br>
      Web Chat
    </td>
    <td>
      Report<br>
      Morning Insights<br>
      Evening Insights<br>
      Global Alert<br>
      Regional Alert<br>
      Pulse Update<br>
      Special Report<br>
      Task Force Report
    </td>
  </tr>
</table>

There is some debate as to whether these uses should reside on the file level or on the intermediate wrapper level. The concern arises from distinct use groups that overlap in terms of files to which they are applicable. For example `DocumentUse` and the proposed `ReportUse` would both apply to documents. Therefore if uses are applied on the file level, the file either have multiple use properties specific to each use group or permit an union of multiple use groups under the use property. Alternately, the two use categories could be applied into a single list of uses, but this would result in a long and unwieldy list. Additionally, certain uses (i.e. pertaining to press materials) are limited to a particular team. Filtering these limited uses out from a larger list would add additional complexity. As such it is probably best to move the use properties to the intermediate wrapper level. This has the added benefit of allow the same document file to be be applied to different uses in different contexts (an admittedly low-probability event).

## Recommendations

1. All non-file specific metadata should be moved off of the `_X_File` types and to an intermediate level wrapper.
1. Following on the above point, create an intermediate wrapper type of `Document`. Migrate existing `DocumentFile` items to a `Document` with a one-to-one relationship with the `DocumentFile`. Once this is done we can remove the non-file specific metadata from the `DocumentFile` instances.
1. Making the above changes would greatly impact the data writing/fetching for the press guidance documents. We would need to rewrite a lot of the queries and components dealing with the publication and rendering of press materials.
1. Maybe - Create an intermediate wrapper type of `Graphic` to wrap the `ImageFile` items in the the graphic projects' images array. This may not be necessary since all of the `ImageFile` properties (with the possible exception of `social`, `style`, `title`, and `use`) are appropriate file-level data.
