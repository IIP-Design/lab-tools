import { Command, flags as cmdFlags } from '@oclif/command';

import { capitalize } from './utils/text';

class GpaLabCreateLabApp extends Command {
  static description =
    'Bootstrap an app using the default GPA/LAB configurations.';

  static flags = {
    version: cmdFlags.version( { 'char': 'v' } ),
    help: cmdFlags.help( { 'char': 'h' } ),
    bundler: cmdFlags.string(
      {
        'char': 'b',
        description:
          'identify the bundler that you would like to use for the project',
        options: ['parcel', 'webpack'],
        required: false,
      },
    ),
    css: cmdFlags.string(
      {
        'char': 'c',
        description:
          'configure CSS type for the app, you can combine multiple options',
        multiple: true,
        options: [
          'css',
          'sass',
          'modules',
        ],
        required: false,
      },
    ),
    typescript: cmdFlags.boolean(
      { 'char': 't', description: 'add TypeScript support to the project' },
    ),
  };

  static args = [
    {
      name: 'project-directory',
      required: true,
      description: 'the name of of your application, we will create.',
      hidden: false,
    },
  ];

  async run() {
    const { args, flags } = this.parse( GpaLabCreateLabApp );

    if ( args['project-directory'] ) {
      this.log( args['project-directory'] );
    }

    if ( flags.bundler ) {
      this.log( `Setting up a ${capitalize( flags.bundler )} project...` );
    }

    if ( flags.typescript ) {
      this.log( 'typescript' );
    }

    return null;
  }
}

export = GpaLabCreateLabApp;
