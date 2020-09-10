import { Command, flags as cmdFlags } from '@oclif/command';

class GpaLabCreateLabApp extends Command {
  static description = 'describe the command here';

  static flags = {
    version: cmdFlags.version( { 'char': 'v' } ),
    help: cmdFlags.help( { 'char': 'h' } ),
    name: cmdFlags.string( { 'char': 'n', description: 'name to print' } ),
    force: cmdFlags.boolean( { 'char': 'f' } ),
  };

  static args = [{ name: 'file' }];

  async run() {
    const { args, flags } = this.parse( GpaLabCreateLabApp );

    const name = flags.name ?? 'world';

    this.log( `hello ${name} from ./src/index.ts` );
    if ( args.file && flags.force ) {
      this.log( `you input --force and --file: ${args.file}` );
    }

    return null;
  }
}

export = GpaLabCreateLabApp;
